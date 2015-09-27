'use strict'

FileResource = ($resource) ->
    $resource '/api/file/:a1/:a2/:a3/:a4/:a5',
      a1: "@a1"
      a2: "@a2"
      a3: "@a3"
      a4: "@a4"
      a5: "@a5"
    ,
      get: {method:'GET'}
      post: {method:'POST'}
      patch: {method:'PATCH'}
      delete: {method:'DELETE'}

angular
    .module('appServices')
    .factory 'FileResource', ['$resource', FileResource]
