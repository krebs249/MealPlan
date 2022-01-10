import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon } from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class Recipes extends Component {
    constructor(props) {
        super(props);
        this.state = {
            recipes: []
        }
    }
    onClick = () => {
        axios
            .get(endpoint + "/api/recipes")
            .then(res => {
                console.log(res.data.recipes)
                this.setState({recipes: JSON.stringify(res.data.recipes)})
            })
            .catch(function (error) {
                console.log(error);
            });

        console.log("test Submit")
    }
    render() {
        console.log("inside render")
        return (
            <div>
                <button onClick={this.onClick}>
                    Click me!
                </button>
                <p>{this.state.recipes}</p>
            </div>
        )
    }
}

export default Recipes;