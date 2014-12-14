define [
  "angular"
  "filters"
  "services"
  "directives"
  "controllers"
  "angularRoute"
], (angular, filters, services, directives, controllers) ->

  # Declare app level module which depends on filters, and services
  angular.module "myApp", [
    "ngRoute"
    "myApp.filters"
    "myApp.services"
    "myApp.directives"
    "myApp.controllers"
  ]
