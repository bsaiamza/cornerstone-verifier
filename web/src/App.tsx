import { lazy, Suspense, useEffect, useMemo, useState } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { CssBaseline, ThemeProvider } from '@mui/material'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'

import { getAppTheme } from './assets/styles/theme'
import { Layout } from './components/Layout'
import { Spinner } from './components/Spinner'
import { ThemeModeContext } from './contexts'
import { DARK_MODE_THEME, LIGHT_MODE_THEME } from './utils/constants'

const Home = lazy(() => import('./pages/Home'))
const Credential = lazy(() => import('./pages/Credential'))
const NotFound = lazy(() => import('./pages/404'))

function App() {
	const [mode, setMode] = useState<typeof LIGHT_MODE_THEME | typeof DARK_MODE_THEME>(LIGHT_MODE_THEME)

	const themeMode = useMemo(
		() => ({
			toggleThemeMode: () => {
				setMode((prevMode) => (prevMode === LIGHT_MODE_THEME ? DARK_MODE_THEME : LIGHT_MODE_THEME))
			},
		}),
		[]
	)

	const theme = useMemo(() => getAppTheme(mode), [mode])

	return (
		<Suspense fallback={<Spinner />}>
			<ThemeModeContext.Provider value={themeMode}>
				<ThemeProvider theme={theme}>
					<CssBaseline />
					<Router>
						<Layout>
							<Routes>
								<Route key='main-route' path='/' element={<Home />} />
								<Route key='credential-route' path='/verify-credential' element={<Credential />} />
								<Route key='404-route' path='*' element={<NotFound />} />
							</Routes>
						</Layout>
					</Router>
				</ThemeProvider>
			</ThemeModeContext.Provider>

			<ToastContainer limit={1} />
		</Suspense>
	)
}

export default App
