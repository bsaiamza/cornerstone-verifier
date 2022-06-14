// components
import BoxComponent from './Box'
import Copyright from './Copyright'

const Footer = () => {
  return (
    <>
      <BoxComponent
        component="footer"
        sx={{
          position: 'fixed',
          bottom: 0,
          width: '100%',
          backgroundColor: '#eee',
          padding: '2rem',
          boxShadow: '0px -1px 0px #ccc',
          zIndex: '2',
        }}
      >
        <Copyright />
      </BoxComponent>
    </>
  )
}

export default Footer
