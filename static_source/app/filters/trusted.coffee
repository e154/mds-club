#
# Created by delta54 on 9/27/15.
#
# http://stackoverflow.com/questions/21292114/external-resource-not-being-loaded-by-angularjs
#

'use strict'

trusted = ($sce) ->
  (url)=>
    $sce.trustAsResourceUrl(url)

angular
  .module('appFilters')
  .filter 'trusted', ['$sce', trusted]
