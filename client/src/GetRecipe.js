import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class Recipes extends Component {
    onClick = () => {
        axios
            .get(endpoint + "/api/recipes")
            .then(res => {
                console.log(res)
            })
            .catch(function (error) {
                console.log(error);
            });

        console.log("test Submit")
    }
    render() {
        return (
            <button onClick={this.onClick}>
                Click me!
            </button>
        )
    }
}

export default Recipes;