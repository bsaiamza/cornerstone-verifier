import { Box } from '@mui/material'

const BoxComponent = ({ alt, children, component, src, sx }) => {
  return (
    <Box alt={alt} component={component} src={src} sx={sx}>
      {children}
    </Box>
  )
}

export default BoxComponent
