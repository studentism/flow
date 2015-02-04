// for conciseness
var $ = go.GraphObject.make;

var flowDiag = $(go.Diagram, "flowDiv",
  {
    initialContentAlignment: go.Spot.Center, // center Diagram contents
    "undoManager.isEnabled": true // enable Ctrl-Z to undo and Ctrl-Y to redo
  });

var myModel = $(go.Model);
// In our model data, each node is represented by a JavaScript object:
myModel.nodeDataArray = [
  { key: "Alpha" },
  { key: "Beta" },
  { key: "Gamma" }
];
myWorkflow.model = myModel;
