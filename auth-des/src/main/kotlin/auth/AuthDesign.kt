package auth

import ee.design.*
import ee.lang.*

object Auth : Comp({ namespace("ee.auth") }) {
    object Auth : Module() {
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

            val sentDisabledConfirmation = propB().meta()
            val sentEnabledConfirmation = propB().meta()
            val disabled = propB().meta()

            val login = command(username, email, password)
            val enable = updateBy(p(disabled) { value(false) })
            val disable = updateBy(p(disabled) { value(true) })

            val sendEnabledConfirmation = command()
            val sendDisabledConfirmation = command()

            object Handler : AggregateHandler({
                defaultState(state {
                    name("Initial")
                    executeAndProduce(create())
                    handle(eventOf(create())).ifTrue(disabled.yes()).to(Disabled)
                    handle(eventOf(create())).ifFalse(disabled.yes()).to(Enabled)
                })
            }) {

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
                    executeAndProduce(sendDisabledConfirmation)
                    handle(eventOf(enable)).to(Enabled).produce(sendEnabledConfirmation)
                })

                object Enabled : State({
                    superUnit(Exist)
                    executeAndProduce(disable)
                    executeAndProduce(sendEnabledConfirmation)
                    handle(eventOf(disable)).to(Disabled).produce(sendDisabledConfirmation)
                    executeAndProduce(delete())
                    handle(eventOf(delete())).to(Deleted)
                })

                object Deleted : State()
            }
        }
    }
}