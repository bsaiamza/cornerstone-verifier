import { Box } from '@mui/material'
import { FingerprintSpinner } from 'react-epic-spinners'

export const Spinner = () => {
	return (
		<Box
			sx={{
				position: 'fixed',
				top: '40%',
				left: '45%',
				backgroundColor: '#08131B',
				borderRadius: 5,
			}}>
			<FingerprintSpinner color='#FAA61A' size={100} />
		</Box>
	)
}
