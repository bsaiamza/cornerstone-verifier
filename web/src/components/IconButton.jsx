import { IconButton } from '@mui/material'

const IconButtonComponent = ({
  ariaControls,
  ariaHaspopup,
  areaLabel,
  children,
  color,
  edge,
  onClick,
  onMouseDown,
  size,
  sx,
}) => {
  return (
    <IconButton
      aria-controls={ariaControls}
      aria-haspopup={ariaHaspopup}
      area-label={areaLabel}
      color={color}
      edge={edge}
      onClick={onClick}
      onMouseDown={onMouseDown}
      size={size}
      sx={sx}
    >
      {children}
    </IconButton>
  )
}

export default IconButtonComponent
