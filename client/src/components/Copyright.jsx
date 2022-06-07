// import { Link } from '@mui/material'
import TypographyComponent from './Typography'

const Copyright = props => {
  return (
    <TypographyComponent
      align="center"
      color="text.secondary"
      variant="body2"
      sx={{ wordWrap: 'break-word', p: 0, m: 0 }}
    >
      {'Copyright © Cornerstone Verifier | '}
      {/* <Link
        color="inherit"
        href="https://www.bankservafrica.com/website/"
        target="_blank"
      >
        BankservAfrica
      </Link>{' '} */}
      {new Date().getFullYear()}
      {'.'}
    </TypographyComponent>
  )
}

export default Copyright
