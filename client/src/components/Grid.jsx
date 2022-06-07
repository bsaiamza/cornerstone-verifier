import { Grid } from '@mui/material'

const GridComponent = ({
  children,
  container,
  item,
  justify,
  md,
  spacing,
  xs,
}) => {
  return (
    <Grid
      container={container}
      item={item}
      justify={justify}
      md={md}
      spacing={spacing}
      xs={xs}
    >
      {children}
    </Grid>
  )
}

export default GridComponent
