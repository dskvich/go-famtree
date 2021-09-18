import React from 'react';
import { render } from 'react-dom';
// Import Highcharts
import Highcharts from "highcharts";
import sankey from "highcharts/modules/sankey.js";
import organization from "highcharts/modules/organization.js";
import HighchartsReact from "highcharts-react-official";

sankey(Highcharts);
organization(Highcharts);

var nodes = [
  {
    id: "1",
    name: "Son",
    description: "7/5/1990",
    image: "https://i.pinimg.com/236x/f1/9c/2f/f19c2f3dc0e1dc6a0a16c7f9138abea1.jpg"
  },
  {
    id: "2",
    name: "Dad",
    column: 1,
    sortIndex: 1,
    description: "7/5/1990",
    image: "https://us.123rf.com/450wm/yupiramos/yupiramos1609/yupiramos160912753/62358474-%EC%95%84%EB%B0%94%ED%83%80-%EB%B9%84%EC%A6%88%EB%8B%88%EC%8A%A4-%EB%82%A8%EC%9E%90-%EC%9B%83-%EB%A7%8C%ED%99%94-%EC%9E%85%EA%B3%A0-%EC%96%91%EB%B3%B5%EA%B3%BC-%EB%84%A5%ED%83%80%EC%9D%B4-%EB%B2%A1%ED%84%B0-%EC%9D%BC%EB%9F%AC%EC%8A%A4%ED%8A%B8-%EB%A0%88%EC%9D%B4-%EC%85%98.jpg"
  },
  {
    id: "3",
    name: "Mom",
    sortIndex: 2,
    column: 1,
    description: "7/5/1990",
    image: "https://cdn5.f-cdn.com/contestentries/1101702/26087492/5999791cc5f03_thumb900.jpg"
  },
  {
    id: "4",
    name: "Grandpa",
    description: "7/5/1990",
    image: "https://thumbs.dreamstime.com/z/old-man-cartoon-icon-happy-over-white-background-colorful-design-illustration-88668344.jpg"
  },
  {
    id: "5",
    name: "Grandma",
    description: "7/5/1990",
    image: "https://t3.ftcdn.net/jpg/02/74/86/30/360_F_274863032_xgwvNFF0u9vZAGPmtvCRHdGxIeeDLAb9.jpg"
  },
  {
    id: "6",
    name: "Grandpa",
    description: "7/5/1990",
    image: "https://st3.depositphotos.com/1007566/12818/v/600/depositphotos_128182398-stock-illustration-grandfather-character-member-avatar.jpg"
  },
  {
    id: "7",
    name: "Grandma",
    description: "7/5/1990",
    image: "https://us.123rf.com/450wm/alexutemov/alexutemov1605/alexutemov160500616/56958889-old-people-cute-granny-and-funny-cute-granny-face-cute-granny-vector-character-and-cartoon-cute-happ.jpg"
  },
  {
    id: "8",
    name: "Uncle",
    sortIndex: 3,
    column: 1,
    description: "7/5/1990",
    image: "https://st2.depositphotos.com/1007566/12304/v/950/depositphotos_123041444-stock-illustration-avatar-man-cartoon.jpg"
  },
  {
    id: "9",
    name: "Aunt",
    column: 1,
    description: "7/5/1990",
    image: "https://us.123rf.com/450wm/azuzl/azuzl1612/azuzl161200010/67501351-belle-fille-blonde-avec-portrait-cheveux-raides-isol%C3%A9-sur-fond-blanc.jpg"
  },
];

var data = [
  ['1', '2'],
  ['1', '3'],
  ['2', '4'],
  ['2', '5'],
  ['3', '6'],
  ['3', '7'],
  ['8', '4'],
  ['8', '5'],
  ['9', '6'],
  ['9', '7'],
];

const options = {
  chart: {
    height: 500,
    inverted: true
  },
  title: {
    text: ""
  },
  series: [{
    type: 'organization',
    keys: ['from', 'to'],

    cursor: 'pointer',
    point: {
        events: {
            click: function () {
                alert('id: ' + this.id);
            }
        }
    },

    data: data,
    linkColor: "#ddd",
    linkLineWidth: 2,
    linkRadius: 60,

    nodes: nodes,


    colorByPoint: false,
    color: '#DBC089',
    dataLabels: {
      color: 'white'
    },
    shadow: {
      color: '#ccc',
      width: 10,
      offsetX: 0,
      offsetY: 0
    },
    animation: false,
    hangingIndent: 10,
    borderColor: '#ccc',
    nodePadding: 20,
    nodeWidth: 80
  }],
  credits: {
    enabled: false
  },
    plotOptions: {
            networkgraph: {
                keys: ['from', 'to'],
                layoutAlgorithm: {
                    enableSimulation: true,
                    friction: -0.9
                }
            }
        },
  tooltip: {
    outside: true,
    formatter: function() {
      return false
    }
  },
  exporting: {
    allowHTML: true,
    sourceWidth: 800,
    sourceHeight: 400
  }

};

class App extends React.Component {
  render() {
    return (
      <div>
        <h1>Family tree</h1>
        <HighchartsReact
          highcharts={Highcharts}
          options={options}
        />
      </div>
    );
  }
}
render(<App />, document.getElementById('root'));
