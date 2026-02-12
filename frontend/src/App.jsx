import './font/pretendard.css'

import style from './App.module.css'

export default function App() {
  // Example word for hangman blanks
  const word = ['', '', '', '', '', '', '', '', '', ''];

  return (
    <>
      <div className={style.main}>
        <div className={style.word_blanks}>
          {word.map((letter, idx) => (
            <div className={style.blank}>
              <p className={style.word}>{letter}</p>
            </div>
          ))}
        </div>
        <input className={style.input_box} type="text" placeholder="🔒 코드 입력" />
        <button className={style.submit_button} type="submit">제출</button>
      </div>
    </>
  )
}