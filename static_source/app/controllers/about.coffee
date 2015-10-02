'use strict'

angular
  .module('appControllers')
  .controller 'aboutCtrl', ['$scope'
  ($scope) ->

    debug = true
    isScan = false
    socket = null

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
        if debug
          console.info("Connection open.")

      socket.onmessage = (e) ->
#        data = JSON.parse(e.data)
        console.log(e)

    ##################################
    # etc
    ##################################

    $scope.startScan = ()->
      if socket
        socket.send(JSON.stringify({key: 'command', value: 'scan start 0 3'}))

    $scope.stopScan = ()->
      if socket
        socket.send(JSON.stringify({key: 'command', value: 'scan stop'}))

    ##################################
    # init
    ##################################
    connect()

  ]