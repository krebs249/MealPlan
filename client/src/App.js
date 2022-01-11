import React from "react";
import "./App.css";
// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";
import Recipes from "./GetRecipe";

function App() {
  return (
    <div>
      <Container>
        <Recipes />
      </Container>
    </div>
  );
}

export default App;