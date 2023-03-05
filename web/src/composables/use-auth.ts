import { reactive, toRefs } from 'vue'
import type { User, LoginCredentials } from './use-auth.types'

const globalState = reactive({
    // --- Composable Initialization State
    initialized: false as boolean,
    init: null as Promise<void> | null,

    // --- User state
    user: null as User | null,
    token: null as string | null,

    // --- User loading state
    loading: false as boolean,
    load: null as Promise<void> | null
})

const useAuth = () => {
    if (!globalState.initialized) {
        globalState.initialized = true
        globalState.init = new Promise((resolve) => {
            const start = Date.now()

            // --- Initialize the composable here

            const end = Date.now()
            console.log(`[useAuth] initialized in ${end - start}ms`)
            setTimeout(() => resolve(), 6*1000)
        })
    }

    function login(creds: LoginCredentials) {}

    return {
        globalState,
        ...toRefs(globalState)
    }
}

export default useAuth
