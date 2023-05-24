import type { Liff } from "@line/liff";
import type { NextPage } from "next";
import Head from "next/head";
import styles from "../styles/Home.module.css";
import { useState, useEffect } from "react"

const Home: NextPage<{ liff: Liff | null; liffError: string | null }> = ({
  liff,
  liffError
}) => {
  const [email, setEmail] = useState()
  const [checkClient, setCheckClient] = useState<Boolean | null>(null)

  useEffect(() => {
    // if (liff) {
    //   console.log(liff.isInClient())
    //   setCheckClient(liff.isInClient())
    // } else {
    //   console.log("nah")
    //   setCheckClient(liff?.isInClient())
    //   // activeBtn()
    // }
    // getUserProfile()
    // if (!liff.isInClient()) {
    //   setCheckClient("Is in client")
    // } else {
    //   setCheckClient("Is in external browser")
    // }
    return () => {
      getUserProfile()
      console.log("asd")
    }
  }, [])

  const getUserProfile = async () => {
    const profile = await liff?.getProfile()
    console.log(profile)
    console.log(liff.getDecodedIDToken().email)
    setEmail(liff.getDecodedIDToken().email)
  }
  return (
    <div>
      <Head>
        <title>LIFF App</title>
        <meta name='viewport' content='width=device-width, initial-scale=1.0' />
        <link rel='icon' href='/favicon.ico' />
      </Head>

      <main className={styles.main}>
        <h1>create-liff-app</h1>
        {liff && <p>LIFF init succeeded. </p>}
        <p>
          <code>{email}</code>
          {liff && checkClient && <p>In client. </p>}
          {liff && !checkClient && <p>Not In client. </p>}
        </p>
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
};

export default Home;
