package auth

import ee.design.gen.angular.DesignAngularGenerator
import ee.design.gen.angular.DesignAngularGeneratorEE
import ee.design.gen.go.DesignGoGenerator
import ee.lang.integ.dPath

fun main() {
    generateGo()
}

fun generateGo() {
    //val generator = DesignGoGenerator(Auth, true)
    //generator.generate(dPath, generator.generatorFactory.goEventDriven())
    val generatorAngular = DesignAngularGeneratorEE(Auth)
    generatorAngular.generate(dPath)
}

