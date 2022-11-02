import React,{ useState} from 'react';
import './App.css';

function App(): JSX.Element {
  const [inputvalue, setInputValue] = useState("");
  const [result, setResult] = useState("");

  const fetchPost = () => {
    const fetchUrl = 'http://localhost:8080/'
    fetch(fetchUrl, {
      method: 'POST',
      body: JSON.stringify({
        message: inputvalue,
      })
    })
      .then((response) => {
        if (response.status === 200) {
          return Promise.resolve(response.json());
        }
        return Promise.reject();
      })
      .then((json) => {
        setResult(json!.message);
      });
      
  };

  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    setInputValue(e.target.value);
  }


  return (
    <div className="App">
      <input type='text' value={inputvalue} onChange={(e) => handleChange(e)} className="inputText"/>
      <br></br>
      <button onClick={() => fetchPost()}>送信</button>
      <p>{result}</p>
    </div>
  );
}

export default App;