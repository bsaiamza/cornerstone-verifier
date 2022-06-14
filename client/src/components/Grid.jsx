import { Grid } from '@mui/material'

const GridComponent = ({
  children,
  container,
  item,
  justify,
  md,
  spacing,
  sx,
  xs,
}) => {
  return (
    <Grid
      container={container}
      item={item}
      justify={justify}
      md={md}
      spacing={spacing}
      sx={sx}
      xs={xs}
    >
      {children}
    </Grid>
  )
}

export default GridComponent
