// react
import { Suspense, lazy } from 'react'
// react-router-dom
import { Routes, Route } from 'react-router-dom'
// mui5 styles
import { ThemeProvider } from '@mui/material/styles'
// react-toastify
import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.min.css'
// context
import { theme } from './context/theme'
// components
const Footer = lazy(() => import('./components/Footer'))
const Navigation = lazy(() => import('./components/Navigation'))
// pages
const ConnectionPage = lazy(() => import('./pages/Connection'))
const VerificationRecordsPage = lazy(() => import('./pages/Records'))
const VerifyCornerstoneCredentialPage = lazy(() =>
  import('./pages/VerifyCornerstone')
)
const VerifyCredentialPage = lazy(() => import('./pages/Verify'))
const HomePage = lazy(() => import('./pages/Home'))
const NotFoundPage = lazy(() => import('./pages/NotFound'))

function App() {
  return (
    <ThemeProvider theme={theme}>
      <Suspense fallback={<div>Loading...</div>}>
        <Navigation />
        <Footer />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="connections" element={<ConnectionPage />} />
          <Route
            path="verify-cornerstone-credential"
            element={<VerifyCornerstoneCredentialPage />}
          />
          <Route path="verify-credential" element={<VerifyCredentialPage />} />
          <Route path="records" element={<VerificationRecordsPage />} />
          <Route path="*" element={<NotFoundPage />} />
        </Routes>

        <ToastContainer
          position="top-center"
          autoClose={5000}
          hideProgressBar={false}
          newestOnTop={false}
          closeOnClick
          limit={3}
          pauseOnFocusLoss
          draggable
          pauseOnHover
        />
      </Suspense>
    </ThemeProvider>
  )
}

export default App
