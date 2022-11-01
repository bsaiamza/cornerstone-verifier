import { useState } from 'react'
import {
	Alert,
	AlertTitle,
	Card,
	CardActionArea,
	CardContent,
	Collapse,
	Grid,
	IconButton,
	Tooltip,
	Typography,
} from '@mui/material'
import { useTheme } from '@mui/material/styles'
import { Close } from '@mui/icons-material'
import { toast } from 'react-toastify'
import axios from 'axios'
import QRCode from 'react-qr-code'

import iamza from '../../assets/images/iamza.png'
import contactable from '../../assets/images/contactable.png'
import address from '../../assets/images/address.png'
import scan from '../../assets/images/scan.png'
import scan_dark from '../../assets/images/scan_dark.png'
import { IAMZA_VERIFIER_URL, LIGHT_MODE_THEME } from '../../utils/constants'

const cornerstoneURL = IAMZA_VERIFIER_URL + '/verify-cornerstone'
const contactableURL = IAMZA_VERIFIER_URL + '/verify-contactable'
const addressURL = IAMZA_VERIFIER_URL + '/verify-address'

const Verify = () => {
	const theme = useTheme()

	const [cornerstoneCredential, setCornerstoneCredential] = useState<string | undefined>('')
	const [contactableCredential, setContactableCredential] = useState<string | undefined>('')
	const [addressCredential, setAddressCredential] = useState<string | undefined>('')
	const [open, setOpen] = useState<boolean | undefined>(true)

	const handleCornerstone = async () => {
		await toast.promise(
			axios
				.get(cornerstoneURL)
				.then((response) => {
					setCornerstoneCredential(response.data.proofRequest)
					toast.success('Generated verification request!')
				})
				.catch((error) => console.log(error)),
			{
				pending: 'Generating...',
			}
		)
	}

	const handleContactable = async () => {
		await toast.promise(
			axios
				.get(contactableURL)
				.then((response) => {
					setContactableCredential(response.data.proofRequest)
					toast.success('Generated verification request!')
				})
				.catch((error) => console.log(error)),
			{
				pending: 'Generating...',
			}
		)
	}

	const handleAddress = async () => {
		await toast.promise(
			axios
				.get(addressURL)
				.then((response) => {
					setAddressCredential(response.data.proofRequest)
					toast.success('Generated verification request!')
				})
				.catch((error) => console.log(error)),
			{
				pending: 'Generating...',
			}
		)
	}

	return (
		<Grid container spacing={1}>
			{/* <Grid item xs={12} md={3} sx={{ borderRight: { xs: 0, md: 1 }, borderBottom: { xs: 1, md: 0 } }}> */}
			<Grid item container xs={12} md={4} direction='column' alignItems='center' justifyContent='center'>
				<div style={{ marginBottom: 10 }}>
					{cornerstoneCredential ? (
						<Collapse in={open}>
							<Alert
								severity='info'
								action={
									<Tooltip title='Close'>
										<IconButton
											aria-label='close'
											color='inherit'
											size='small'
											onClick={() => {
												setOpen(false)
											}}>
											<Close fontSize='inherit' />
										</IconButton>
									</Tooltip>
								}
								sx={{ mb: 2, textAlign: 'left' }}>
								<AlertTitle>Click on the credential again to refresh QR code.</AlertTitle>
							</Alert>
						</Collapse>
					) : (
						''
					)}
				</div>
				<Card
					sx={{
						height: 220,
						width: 220,
						backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '#fff' : '',
						borderRadius: 5,
					}}
					elevation={3}>
					<CardActionArea sx={{ textAlign: 'center', p: 1 }} onClick={handleCornerstone}>
						<img src={iamza} height='100' width='100' alt='iamza' />
						<CardContent>
							<Typography gutterBottom variant='subtitle2' component='div' color='text.secondary'>
								Cornerstone Credential
							</Typography>
							<Typography gutterBottom variant='caption' component='div' color='text.secondary'>
								Verify your age
								<br />
								<span style={{ fontSize: 20 }}>🔞</span>
							</Typography>
						</CardContent>
					</CardActionArea>
				</Card>
				<div style={{ marginTop: 10 }}>
					{cornerstoneCredential ? (
						<>
							{theme.palette.mode === LIGHT_MODE_THEME ? (
								<img src={scan} height='200' width='200' alt='scan' />
							) : (
								<img src={scan_dark} height='200' width='200' alt='scan' />
							)}
							<div
								style={{
									backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '' : '#F5F5F5',
									padding: theme.palette.mode === LIGHT_MODE_THEME ? '' : 5,
									borderRadius: theme.palette.mode === LIGHT_MODE_THEME ? '' : 5,
								}}>
								<QRCode value={cornerstoneCredential} />
							</div>
						</>
					) : (
						''
					)}
				</div>
			</Grid>
			<Grid item container xs={12} md={4} direction='column' alignItems='center' justifyContent='center'>
				<div style={{ marginBottom: 10 }}>
					{contactableCredential ? (
						<Collapse in={open}>
							<Alert
								severity='info'
								action={
									<Tooltip title='Close'>
										<IconButton
											aria-label='close'
											color='inherit'
											size='small'
											onClick={() => {
												setOpen(false)
											}}>
											<Close fontSize='inherit' />
										</IconButton>
									</Tooltip>
								}
								sx={{ mb: 2, textAlign: 'left' }}>
								<AlertTitle>Click on the credential again to refresh QR code.</AlertTitle>
							</Alert>
						</Collapse>
					) : (
						''
					)}
				</div>
				<Card
					sx={{
						height: 220,
						width: 220,
						backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '#fff' : '',
						borderRadius: 5,
					}}
					elevation={3}>
					<CardActionArea sx={{ textAlign: 'center', p: 1 }} onClick={handleContactable}>
						<img src={contactable} height='100' width='100' alt='contactable' />
						<CardContent>
							<Typography gutterBottom variant='subtitle2' component='div' color='text.secondary'>
								Contactable Credential
							</Typography>
							<Typography gutterBottom variant='caption' component='div' color='text.secondary'>
								Verify your identity attributes are valid
							</Typography>
						</CardContent>
					</CardActionArea>
				</Card>
				<div style={{ textAlign: 'center', marginTop: 10 }}>
					{contactableCredential ? (
						<>
							{theme.palette.mode === LIGHT_MODE_THEME ? (
								<img src={scan} height='200' width='200' alt='scan' />
							) : (
								<img src={scan_dark} height='200' width='200' alt='scan' />
							)}
							<div
								style={{
									backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '' : '#F5F5F5',
									padding: theme.palette.mode === LIGHT_MODE_THEME ? '' : 5,
									borderRadius: theme.palette.mode === LIGHT_MODE_THEME ? '' : 5,
								}}>
								<QRCode value={contactableCredential} />
							</div>
						</>
					) : (
						''
					)}
				</div>
			</Grid>
			<Grid item container xs={12} md={4} direction='column' alignItems='center' justifyContent='center'>
				<div style={{ marginBottom: 10 }}>
					{addressCredential ? (
						<Collapse in={open}>
							<Alert
								severity='info'
								action={
									<Tooltip title='Close'>
										<IconButton
											aria-label='close'
											color='inherit'
											size='small'
											onClick={() => {
												setOpen(false)
											}}>
											<Close fontSize='inherit' />
										</IconButton>
									</Tooltip>
								}
								sx={{ mb: 2, textAlign: 'left' }}>
								<AlertTitle>Click on the credential again to refresh QR code.</AlertTitle>
							</Alert>
						</Collapse>
					) : (
						''
					)}
				</div>
				<Card
					sx={{
						height: 220,
						width: 220,
						backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '#fff' : '',
						borderRadius: 5,
					}}
					elevation={3}>
					<CardActionArea sx={{ textAlign: 'center', p: 1 }} onClick={handleAddress}>
						<img src={address} height='100' width='100' alt='contactable' />
						<CardContent>
							<Typography gutterBottom variant='subtitle2' component='div' color='text.secondary'>
								Physical Address Credential
							</Typography>
							<Typography gutterBottom variant='caption' component='div' color='text.secondary'>
								Verify your address is valid
							</Typography>
						</CardContent>
					</CardActionArea>
				</Card>
				<div style={{ textAlign: 'center', marginTop: 10 }}>
					{addressCredential ? (
						<>
							{theme.palette.mode === LIGHT_MODE_THEME ? (
								<img src={scan} height='200' width='200' alt='scan' />
							) : (
								<img src={scan_dark} height='200' width='200' alt='scan' />
							)}
							<div
								style={{
									backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '' : '#F5F5F5',
									padding: theme.palette.mode === LIGHT_MODE_THEME ? '' : 5,
									borderRadius: theme.palette.mode === LIGHT_MODE_THEME ? '' : 5,
								}}>
								<QRCode value={addressCredential} />
							</div>
						</>
					) : (
						''
					)}
				</div>
			</Grid>
		</Grid>
	)
}

export default Verify
