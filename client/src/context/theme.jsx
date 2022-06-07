import { createTheme } from '@mui/material/styles'

export const theme = createTheme({
  palette: {
    primary: {
      main: '#faa61a',
      light: '#faa61a41',
    },
    white: {
      main: '#ffffff',
    },
    black: {
      main: '#000000',
    },
    grey: {
      main: '#000000',
    },
    dark: {
      main: '#08131B',
    },
    orange: {
      main: '#F5B344',
    },
  },
  components: {
    MuiDrawer: {
      styleOverrides: {
        paper: {
          background: '#eee',
          width: 250,
          padding: '1rem',
        },
      },
    },
    MuiAppBar: {
      styleOverrides: {
        root: {
          background: '#08131B',
        },
      },
    },
  },
})
