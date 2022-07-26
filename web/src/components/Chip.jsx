// mui5
import { Chip } from '@mui/material'

const ChipComponent = ({ avatar, color, icon, label, size, sx, variant }) => {
  return (
    <Chip
      avatar={avatar}
      color={color}
      icon={icon}
      label={label}
      size={size}
      sx={sx}
      variant={variant}
    />
  )
}

export default ChipComponent
