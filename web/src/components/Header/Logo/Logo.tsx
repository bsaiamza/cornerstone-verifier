import { NavLink } from 'react-router-dom'
import { styled } from '@mui/material'
import { useTheme } from '@mui/material/styles'

import { LIGHT_MODE_THEME } from '../../../utils/constants'
import LogoLight from '../../../assets/images/logo_light.png'
import LogoDark from '../../../assets/images/logo_dark.png'

export const Logo = () => {
	const theme = useTheme()

	return (
		<>
			{theme.palette.mode === LIGHT_MODE_THEME ? (
				<StyledNavLink to='/'>
					<img src={LogoLight} alt='IAMZA' height={100} />
				</StyledNavLink>
			) : (
				<StyledNavLink to='/'>
					<img src={LogoDark} alt='IAMZA' height={100} />
				</StyledNavLink>
			)}
		</>
	)
}

const StyledNavLink = styled(NavLink)`
	text-decoration: none;
	color: inherit;
`
