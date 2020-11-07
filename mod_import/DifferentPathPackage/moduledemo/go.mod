module Go_exercise/mod_import/DifferentPathPackage/moduledemo

go 1.15

require "tempconv" v0.0.0
replace "tempconv" => "../mypackage/tempconv"
