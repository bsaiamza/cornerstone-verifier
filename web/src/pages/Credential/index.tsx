import { SetStateAction, useEffect, useState } from 'react'
import { Grid, Tab, Tabs, Typography } from '@mui/material'

import { SEO } from '../../components/SEO'
import { TabPanel } from '../../components/Tabs'
import Verify from './Verify'
import VerifyByEmail from './VerifyByEmail'
import { Dvr, ForwardToInbox } from '@mui/icons-material'

const a11yProps = (index: any) => {
	return {
		id: `custom-tab-${index}`,
		'aria-controls': `custom-tabpanel-${index}`,
	}
}

const breakpoints = {
	values: {
		xs: 0,
		sm: 576,
		md: 768,
		lg: 992,
		xl: 1200,
		xxl: 1400,
	},
}

const Credential = () => {
	const [tabsOrientation, setTabsOrientation] = useState('horizontal')
	const [tabValue, setTabValue] = useState(0)

	useEffect(() => {
		function handleTabsOrientation() {
			return window.innerWidth < breakpoints.values.sm
				? setTabsOrientation('vertical')
				: setTabsOrientation('horizontal')
		}

		window.addEventListener('resize', handleTabsOrientation)

		handleTabsOrientation()

		return () => window.removeEventListener('resize', handleTabsOrientation)
	}, [tabsOrientation])

	const handleSetTabValue = (event: any, newValue: SetStateAction<number>) => setTabValue(newValue)

	return (
		<>
			<SEO title='Verify' />

			<div style={{ textAlign: 'center', paddingBottom: 10, alignItems: 'center' }}>
				<Typography gutterBottom variant='h4' component='div' color='text.primary'>
					Verifiable Credentials
				</Typography>
				<Grid container direction='column' alignItems='center' justifyContent='center'>
					<Tabs
						// @ts-ignore
						orientation={tabsOrientation}
						value={tabValue}
						onChange={handleSetTabValue}
						sx={{ width: { lg: '30%' } }}>
						<Tab label='Display' icon={<Dvr />} {...a11yProps(0)} />
						<Tab label='Email' icon={<ForwardToInbox />} {...a11yProps(1)} />
					</Tabs>
				</Grid>
			</div>
			<TabPanel value={tabValue} index={0}>
				<Verify />
			</TabPanel>
			<TabPanel value={tabValue} index={1}>
				<VerifyByEmail />
			</TabPanel>
		</>
	)
}

export default Credential
