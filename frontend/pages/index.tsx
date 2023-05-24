import type { Liff } from "@line/liff"
import type { NextPage } from "next"
import Head from "next/head"
import styles from "../styles/Home.module.css"
import { useState, useEffect } from "react"

const Home: NextPage<{ liff: Liff | null; liffError: string | null }> = ({
  liff,
  liffError,
}) => {
  const [email, setEmail] = useState()
  const [checkClient, setCheckClient] = useState<Boolean | null>(null)
  let userProfile: any | null = null

  useEffect(() => {
    // Use the liff object or handle the liffError here
    if (liff) {
      console.log("fuxk")
      getUserProfile()
    } else if (liffError) {
      // There was an error initializing liff
      // Handle the error
      console.log("fuxk999999")
    }
  }, [liff, liffError])

  const getUserProfile = async () => {
    const profile = await liff.getProfile()
    console.log(profile)
    console.log(liff.getDecodedIDToken().email)
    setEmail(liff.getDecodedIDToken().email)
  }

  // if (liff) {
  //   getUserProfile()
  //   // You can use the userProfile object or perform any desired operations with it here
  // }

  return (
    <div>
      <Head>
        <title>LIFF App</title>
        <meta name='viewport' content='width=device-width, initial-scale=1.0' />
        <link rel='icon' href='/favicon.ico' />
      </Head>

      <main className={styles.main}>
        <h1>create-liff-app</h1>
        {liff && <p>{email}</p>}
        {liff && !checkClient && <p>Not In client. </p>}
        {liffError && (
          <>
            <p>LIFF init failed.</p>
            <p>
              <code>{liffError}</code>
            </p>
          </>
        )}
        <a
          href='https://developers.line.biz/ja/docs/liff/'
          target='_blank'
          rel='noreferrer'
        >
          LIFF Documentation
        </a>
      </main>
    </div>
  )
}

export default Home
