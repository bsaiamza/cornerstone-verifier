import React, { useState } from 'react'
import { AppBar, Box, Toolbar } from '@mui/material'
import { useTheme } from '@mui/material/styles'

import { ThemeSwitcher } from './ThemeSwitcher'
import { HEADER_HEIGHT, LIGHT_MODE_THEME } from '../../utils/constants'
import { Logo } from './Logo'

export const Header = () => {
	const theme = useTheme()

	return (
		<>
			<AppBar
				position='fixed'
				sx={{
					height: { HEADER_HEIGHT },
					justifyContent: 'center',
					backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '#fff' : '#08131B',
				}}
				elevation={0}>
				<Toolbar disableGutters variant='dense'>
					<Logo />
					<Box sx={{ flexGrow: 1 }} />
					<Box sx={{ alignItems: 'center' }}>
						<ThemeSwitcher />
					</Box>
				</Toolbar>
			</AppBar>
		</>
	)
}
