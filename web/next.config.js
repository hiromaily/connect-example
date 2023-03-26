/** @type {import('next').NextConfig} */

// const csp = `
//   default-src 'self';
//   script-src 'self';
//   child-src example.com;
//   style-src 'self' localhost:3000;
//   font-src 'self';
// `
const csp = `
  default-src 'self';
  script-src 'self';
  child-src 'none';
  style-src 'unsafe-inline';
  font-src 'unsafe-inline';
`

const headers = [
  {
    key: 'X-Content-Type-Options',
    value: 'nosniff'
  },
  {
    key: 'Referrer-Policy',
    value: 'origin-when-cross-origin'
  },
  // {
  //   key: 'Content-Type',
  //   value: 'text/html; charset=UTF-8'
  // },
  {
    key: 'Strict-Transport-Security',
    value: 'max-age=63072000; includeSubDomains; preload'
  },
  // {
  //   key: 'Access-Control-Allow-Origin',
  //   value: 'https://localhost:8080'
  // },
  // {
  //   key: 'Content-Security-Policy',
  //   value: csp.replace(/\s{2,}/g, ' ').trim()
  // }
]

const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  async headers() {
    return [
      {
        // Apply these headers to all routes in your application.
        source: '/:path*',
        headers: headers,
      },
    ]
  },
}

module.exports = nextConfig
