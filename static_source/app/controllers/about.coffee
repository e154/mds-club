'use strict'

angular
  .module('appControllers')
  .controller 'aboutCtrl', ['$scope', '$timeout'
  ($scope, $timeout) ->

    debug = false
#    launched|stopped
    $scope.status = "stopped"
    $scope.info = ""
    $scope.total = 0
    $scope.current = 0
    $scope.progress = 0
    socket = null
    start = 0
    stop = 0

    ##################################
    # socket
    ##################################
    connect = ()->
      if !("WebSocket") in window
        return

      protocol = if document.location.protocol == "https:" then "wss:" else "ws:"
      url = window.location.href.split('/')[2]
      socket = new WebSocket(protocol+"//"+url+"/ws")

      if socket
        $scope.$on '$locationChangeStart', ()=>
          socket.close()

      socket.onclose = (e) ->
        if debug
          console.info("Connection closed.");

      socket.onopen = () ->
        socket.send(JSON.stringify({key: 'command', value: 'scan status'}))
        if debug
          console.info("Connection open.")

      socket.onmessage = (e) ->
        data = angular.fromJson(e.data)

#        status
        if data?.scan_status?.status
          $timeout ()->
            $scope.$apply(
              $scope.status = data.scan_status.status
            )
          , 0

          if debug
            console.log $scope.status

#        info chanel
        if data?.scan_status?.info
          $scope.info = data.scan_status.info
          if debug
            console.log $scope.info

#         total
        if data?.scan_status?.total
          $scope.total = data.scan_status.total
          if debug
            console.log $scope.total
          updateProgress()

#        current
        if data?.scan_status?.current
          $scope.current = data.scan_status.current
          if debug
            console.log $scope.current
          updateProgress()

    ##################################
    # etc
    ##################################

    $scope.startScan = ()->
      if socket
        socket.send(JSON.stringify({key: 'command', value: 'scan start '+start+' '+ stop}))

    $scope.stopScan = ()->
      if socket
        socket.send(JSON.stringify({key: 'command', value: 'scan stop'}))

    updateProgress = ()=>
      $timeout ()=>
        $scope.$apply(
          $scope.progress = Math.ceil($scope.current/Math.floor($scope.total/100))
        )
      , 0

    ##################################
    # init
    ##################################
    connect()

  ]