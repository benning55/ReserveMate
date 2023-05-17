import type { Liff } from "@line/liff"
import type { NextPage } from "next"
import Head from "next/head"
import styles from "../styles/Home.module.css"
import { useState, useEffect } from "react"

const About: NextPage<{ liff: Liff | null; liffError: string | null }> = ({
  liff,
  liffError,
}) => {
  const [checkClient, setCheckClient] = useState<Boolean | null>(null)

  useEffect(() => {
    if (liff) {
      console.log(liff.isInClient())
      setCheckClient(liff.isInClient())
    } else {
      console.log("nah")
      setCheckClient(liff?.isInClient())
      // activeBtn()
    }
    // if (!liff.isInClient()) {
    //   setCheckClient("Is in client")
    // } else {
    //   setCheckClient("Is in external browser")
    // }
  }, [])

  const activeBtn = () => {
    console.log(liff?.isInClient())
    setCheckClient(liff?.isInClient())
  }

  return (
    <div>
      <main className={styles.main}>
        <h1>This is about page</h1>
        {liff && !checkClient && <p>Not in client. </p>}
        {/* <button onClick={activeBtn}>Benning</button> */}
      </main>
    </div>
  )
}

export default About
