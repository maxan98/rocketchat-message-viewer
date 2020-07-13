import React, {useEffect, useState} from 'react';
import ReactDOM from 'react-dom';
import {render} from "react-dom";

function List(){
  return(
      <div>
      <MyComponent />
      </div>
  )
}
function Myh(){
  return(

      <h1>Some text</h1>
  )
}

function MyComponent() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState([]);

    // Примечание: пустой массив зависимостей [] означает, что
    // этот useEffect будет запущен один раз
    // аналогично componentDidMount()
    useEffect(() => {
        fetch("http://localhost:8080/api/v1/5DNt6wt7wJR97A6SZmsyxCiwpsoukWmvEF")
            .then(res => res.json())
            .then(
                (result) => {
                    setIsLoaded(true);
                    setItems(result);
                },
                // Примечание: Обрабатывать ошибки необходимо именно здесь
                // вместо блока catch(), чтобы не пропустить
                // исключения из реальных ошибок в компонентах.
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            )
    }, [])

    if (error) {
        return <div>Ошибка: {error.message}</div>;
    } else if (!isLoaded) {
        return <div>Загрузка...</div>;
    } else {
        return (
            <ul>
                {items.map(item => (
                    <li key={item.Id}>
                        <h2>{item.Rid}</h2> <br /> <h4> {item.User.Username}: {item.Msg} </h4> <br />
                    </li>
                ))}
            </ul>
        );
    }
}
render(<List />, document.getElementById("root"))