import React, { FC } from 'react'

interface SEOProps {
	title: string
}

export const SEO = ({ title }: SEOProps) => {
	return (
		<>
			<meta charSet='utf-8' />
			<meta name='robots' content='noindex, follow' />
			<meta name='description' content='Debi Check Query System' />
			<meta name='viewport' content='width=device-width, initial-scale=1, shrink-to-fit=no' />
			<title>IAMZA | {title}</title>
		</>
	)
}
