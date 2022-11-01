import * as React from 'react'
import { Box } from '@mui/material'

interface TabPanelProps {
	children?: React.ReactNode
	index: number
	value: number
}

export function TabPanel(props: TabPanelProps) {
	const { children, value, index, ...rest } = props
	return (
		<div
			role='tabpanel'
			hidden={value !== index}
			id={`custom-tabpanel-${index}`}
			aria-labelledby={`custom-tab-${index}`}
			// eslint-disable-next-line react/jsx-props-no-spreading
			{...rest}>
			{value === index && <Box sx={{ p: 3 }}>{children}</Box>}
		</div>
	)
}
