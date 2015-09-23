'use strict'

BooksResource = ($resource) ->
    $resource '/api/books/:a1/:a2/:a3/:a4/:a5',
      a1:"@a1"
      a2:"@a2"
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
    .factory 'BooksResource', ['$resource', BooksResource]
