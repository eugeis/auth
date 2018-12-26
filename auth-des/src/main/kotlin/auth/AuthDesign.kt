package auth

import ee.design.*
import ee.lang.*

object Auth : Module({ namespace("auth") }) {
    object PersonName : Basic() {
        val first = propS()
        val last = propS()
    }

    object UserCredentials : Values() {
        val username = propS()
        val password = propS()
    }

    object Account : Entity() {
        val name = prop(PersonName)
        val username = propS().unique()
        val password = propS().hidden()
        val email = propS().unique()
        val roles = propListT(n.String)

        val disabled = propB().meta()

        val login = command(username, email, password)
        val enable = updateBy(p(disabled) { value(false) })
        val disable = updateBy(p(disabled) { value(true) })

        val sendCreatedConfirmation = command()
        val sendEnabledConfirmation = command()
        val sendDisabledConfirmation = command()

        object Handler : AggregateHandler() {
            object Initial : State({
                executeAndProduce(create())
                handle(eventOf(create())).ifTrue(disabled.yes()).to(Disabled)
                handle(eventOf(create())).ifFalse(disabled.yes()).to(Enabled)
            })

            object Exist : State({
                virtual()
                executeAndProduce(update())
                handle(eventOf(update()))

                executeAndProduce(delete())
                handle(eventOf(delete())).to(Deleted)
            })

            object Disabled : State({
                superUnit(Exist)
                executeAndProduce(enable)
                handle(eventOf(enable)).to(Enabled)
            })

            object Enabled : State({
                superUnit(Exist)
                executeAndProduce(disable)
                handle(eventOf(disable)).to(Disabled)
            })

            object Deleted : State()
        }

        object AccountConfirmation : ProcessManager() {
            object Initial : State({
                executeAndProduce(create())
                handle(eventOf(create())).ifTrue(disabled.yes()).to(Disabled)
                handle(eventOf(create())).ifFalse(disabled.yes()).to(Enabled)
            })

            object Disabled : State({
                handle(eventOf(enable)).to(Enabled).produce(sendEnabledConfirmation)
            })

            object Enabled : State({
                handle(eventOf(disable)).to(Disabled).produce(sendDisabledConfirmation)
            })
        }
    }
}