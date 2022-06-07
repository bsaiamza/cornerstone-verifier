import { Divider } from '@mui/material'
// components
import BoxComponent from '../../components/Box'
// import CardComponent from '../../components/Card'
import GridComponent from '../../components/Grid'
import TypographyComponent from '../../components/Typography'
// images
import bg from '../../assets/images/bg1.png'

const HomePage = () => {
  return (
    <BoxComponent sx={{ m: '2rem' }}>
      <GridComponent container spacing={2}>
        <GridComponent item xs={12} md={6}>
          <BoxComponent sx={{ textAlign: 'center' }}>
            <TypographyComponent variant="h2" color="Grey">
              Iamza Cornerstone Verifier
            </TypographyComponent>
            <TypographyComponent
              variant="h5"
              color="Grey"
              sx={{ marginBottom: '1rem' }}
            >
              Identity credentials made reliable and easy.
            </TypographyComponent>
            <Divider />
          </BoxComponent>
        </GridComponent>
        <GridComponent item xs={12} md={6}>
          <BoxComponent
            alt="DI"
            component="img"
            src={bg}
            sx={{
              height: 'auto',
              width: '100%',
              opacity: 0.5,
            }}
          />
        </GridComponent>
      </GridComponent>
    </BoxComponent>
  )
}

export default HomePage
