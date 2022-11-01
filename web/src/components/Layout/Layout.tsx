import { ReactNode } from 'react'
import { Box, Grid, styled, Paper } from '@mui/material'

import { Footer } from '../Footer'
import { FOOTER_HEIGHT, HEADER_HEIGHT } from '../../utils/constants'
import { Header } from '../Header'

interface LayoutProps {
	children: ReactNode
}

export const Layout = ({ children }: LayoutProps) => {
	return (
		<LayoutWrapper>
			<ContentWrapper>
				<Box component='header'>
					<Header />
				</Box>
				<Box component='main' sx={{ flexGrow: 1, p: 5, mt: `calc(${HEADER_HEIGHT}px)` }}>
					<Paper sx={{ borderRadius: 5, height: '100%', p: 5 }}>{children}</Paper>
				</Box>
			</ContentWrapper>
			<Box component='footer'>
				<Footer />
			</Box>
		</LayoutWrapper>
	)
}

const LayoutWrapper = styled('div')`
	min-height: 100vh;
`
const ContentWrapper = styled('div')`
	display: flex;
	min-height: calc(100vh - ${FOOTER_HEIGHT}px);
`
