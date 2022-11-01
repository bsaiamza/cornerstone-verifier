import { createTheme, responsiveFontSizes } from '@mui/material'
import chroma from 'chroma-js'

// constants
import { DARK_MODE_THEME, LIGHT_MODE_THEME } from '../../utils/constants'

function pxToRem(number: any, baseNumber = 16) {
	return `${number / baseNumber}rem`
}

function hexToRgb(color: any) {
	return chroma(color).rgb().join(', ')
}

function rgba(color: any, opacity: any) {
	return `rgba(${hexToRgb(color)}, ${opacity})`
}

function boxShadow(offset = [], radius = [], color: string, opacity: number, inset = '') {
	const [x, y] = offset
	const [blur, spread] = radius

	return `${inset} ${pxToRem(x)} ${pxToRem(y)} ${pxToRem(blur)} ${pxToRem(spread)} ${rgba(color, opacity)}`
}

export const getAppTheme = (mode: typeof LIGHT_MODE_THEME | typeof DARK_MODE_THEME) => {
	let theme = createTheme({
		palette: {
			mode,
			...(mode === 'light'
				? {
						// palette values for light mode
						primary: {
							light: '#F5F5F5',
							main: '#FAA61A',
							dark: '#4D4D4D',
							contrastText: '#ffffff',
						},
						divider: '#333333',
						text: {
							primary: '#333333',
							secondary: '#333333',
						},
						background: {
							default: '#fff',
							paper: '#F5F5F5',
						},
				  }
				: {
						// palette values for dark mode
						primary: {
							light: '#F5F5F5',
							main: '#FAA61A',
							dark: '#4D4D4D',
							contrastText: '#ffffff',
						},
						divider: '#333333',
						background: {
							default: '#08131B',
							paper: '#08131B',
						},
						text: {
							primary: '#ffffff',
							secondary: '#ffffff',
						},
				  }),
		},
		components: {
			MuiTabs: {
				styleOverrides: {
					root: {
						position: 'relative',
						// backgroundColor: '#f8f9fa',
						borderRadius: pxToRem(12),
						minHeight: 'unset',
						padding: pxToRem(4),
					},

					flexContainer: {
						height: '100%',
						position: 'relative',
						zIndex: 10,
					},

					// @ts-ignore
					fixed: {
						overflow: 'unset !important',
						overflowX: 'unset !important',
					},

					vertical: {
						'& .MuiTabs-indicator': {
							width: '100%',
						},
					},

					indicator: {
						height: '100%',
						borderRadius: pxToRem(8),
						backgroundColor: 'inherit',
						// @ts-ignore
						boxShadow: boxShadow([0, 1], [5, 1], '#ddd', 1),
						transition: 'all 500ms ease',
					},
				},
			},
			MuiTab: {
				styleOverrides: {
					root: {
						display: 'flex',
						alignItems: 'center',
						flexDirection: 'row',
						flex: '1 1 auto',
						textAlign: 'center',
						maxWidth: 'unset !important',
						minWidth: 'unset !important',
						minHeight: 'unset !important',
						fontSize: pxToRem(16),
						fontWeight: 400,
						textTransform: 'none',
						lineHeight: 'inherit',
						padding: pxToRem(4),
						borderRadius: pxToRem(8),
						color: `inherit !important`,
						opacity: '1 !important',

						'& .material-icons, .material-icons-round': {
							marginBottom: '0 !important',
							marginRight: pxToRem(6),
						},

						'& svg': {
							marginBottom: '0 !important',
							marginRight: pxToRem(6),
						},
					},

					labelIcon: {
						paddingTop: pxToRem(4),
					},
				},
			},
		},
	})
	theme = responsiveFontSizes(theme)
	return theme
}
