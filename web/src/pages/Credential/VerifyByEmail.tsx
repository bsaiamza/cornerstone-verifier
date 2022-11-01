import { Button, FormControl, Grid, InputLabel, MenuItem, Paper, Select, TextField } from '@mui/material'
import { useTheme } from '@mui/material/styles'
import axios from 'axios'
import { Formik, Form } from 'formik'
import { useState } from 'react'
import { toast } from 'react-toastify'

import { IAMZA_VERIFIER_URL, LIGHT_MODE_THEME } from '../../utils/constants'

const addressURL = IAMZA_VERIFIER_URL + '/verify-address-email'
const cornerstoneURL = IAMZA_VERIFIER_URL + '/verify-cornerstone-email'
const contactableURL = IAMZA_VERIFIER_URL + '/verify-contactable-email'

const VerifyByEmail = () => {
	const theme = useTheme()
	const [submitting, setSubmitting] = useState<boolean | undefined>(false)
	const [credential, setCredential] = useState('')

	const addressEmail = async (values: any) => {
		setSubmitting(true)

		await toast.promise(
			axios
				.post(addressURL, values)
				.then((response: any) => {
					toast.success('Emailed proof request!')
				})
				.catch((error: any) => {
					toast.error(error.response.data.msg)
				}),
			{
				pending: 'Emailing proof request...',
			}
		)
		setSubmitting(false)
	}

	const cornerstoneEmail = async (values: any) => {
		setSubmitting(true)

		await toast.promise(
			axios
				.post(cornerstoneURL, values)
				.then((response: any) => {
					toast.success('Emailed proof request!')
				})
				.catch((error: any) => {
					toast.error(error.response.data.msg)
				}),
			{
				pending: 'Emailing proof request...',
			}
		)
		setSubmitting(false)
	}

	const contactableEmail = async (values: any) => {
		setSubmitting(true)

		await toast.promise(
			axios
				.post(contactableURL, values)
				.then((response: any) => {
					toast.success('Emailed proof request!')
				})
				.catch((error: any) => {
					toast.error(error.response.data.msg)
				}),
			{
				pending: 'Emailing proof request...',
			}
		)
		setSubmitting(false)
	}

	return (
		<Grid container>
			<Grid item xs={0} md={3} />
			<Grid item container xs={12} md={6} direction='column' alignItems='center' justifyContent='center'>
				<Paper
					square
					elevation={2}
					sx={{
						p: 5,
						width: { md: '100%' },
						backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '#fff' : '',
						textAlign: 'center',
						borderRadius: 5,
					}}>
					<Formik
						initialValues={{
							email: '',
						}}
						// validate={idValidation}
						onSubmit={(values) => {
							credential === 'address'
								? addressEmail(values)
								: '' || credential === 'cornerstone'
								? cornerstoneEmail(values)
								: '' || credential === 'contactable'
								? contactableEmail(values)
								: ''
						}}>
						{({ values, handleChange }) => (
							<Form>
								<div>
									<FormControl sx={{ width: '16.5rem' }}>
										<InputLabel id='self_attested' sx={{ margin: '1rem 0 0 1rem' }} required>
											Choose Credential
										</InputLabel>
										<Select
											labelId='credential'
											id='credential'
											name='credential'
											label='Choose Credential'
											value={credential}
											onChange={(e) => setCredential(e.target.value)}
											required
											sx={{ m: '1rem' }}>
											<MenuItem value='address'>Address</MenuItem>
											<MenuItem value='contactable'>Contactable</MenuItem>
											<MenuItem value='cornerstone'>Cornerstone</MenuItem>
										</Select>
									</FormControl>
								</div>
								<div>
									<TextField
										id='email'
										name='email'
										type='email'
										value={values.email}
										onChange={handleChange}
										label='Email'
										sx={{ m: '1rem', width: { md: '50%' } }}
										required
									/>
								</div>

								<div>
									<Button
										variant='contained'
										size='small'
										type='submit'
										sx={{ color: '#fff', m: '1rem' }}
										disabled={submitting}>
										Submit
									</Button>
								</div>
							</Form>
						)}
					</Formik>
				</Paper>
			</Grid>
			<Grid item xs={0} md={3} />
		</Grid>
	)
}

export default VerifyByEmail
