package auth

import ee.design.gen.go.DesignGoGenerator
import ee.lang.integ.dPath

fun main(args: Array<String>) {
    generateGo()
}

fun generateGo() {
    var generator = DesignGoGenerator(Auth)
    generator.generate(dPath)
}

