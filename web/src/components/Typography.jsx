import { Typography } from '@mui/material'

const TypographyComponent = ({ align, children, color, sx, variant }) => {
  return (
    <Typography
      align={align}
      color={color ? color : '#777'}
      sx={sx}
      variant={variant}
    >
      {children}
    </Typography>
  )
}

export default TypographyComponent
