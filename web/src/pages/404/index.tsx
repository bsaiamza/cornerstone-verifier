import React, { ReactElement, FC } from 'react';
import { Box, Typography } from '@mui/material';

// 404 img
import img404 from '../../assets/images/Error404.png';

const NotFound: FC<any> = (): ReactElement => {
	return (
		<Box
			sx={{
				flexGrow: 1,
				display: 'flex',
				justifyContent: 'center',
				alignItems: 'center',
			}}>
			<img src={img404} alt='404' />
		</Box>
	);
};

export default NotFound;
