// mui5
import { Button } from '@mui/material'

const ButtonComponent = ({
  children,
  color,
  disabled,
  endIcon,
  onClick,
  size,
  startIcon,
  sx,
  type,
  variant,
}) => {
  return (
    <Button
      color={color}
      disabled={disabled}
      endIcon={endIcon}
      onClick={onClick}
      size={size}
      sx={sx}
      startIcon={startIcon}
      type={type}
      variant={variant}
    >
      {children}
    </Button>
  )
}

export default ButtonComponent
