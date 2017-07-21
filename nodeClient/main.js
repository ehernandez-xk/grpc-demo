const grpc = require('grpc');
const argv = require('minimist')(process.argv.slice(2));

const PROTO_PATH = __dirname + '/service.proto';
const pb = grpc.load(PROTO_PATH);

var client;
var person = {
    name: ''
};

const port = ':50051';

function addPerson(name) {
    if (!name) {
        console.log('Parameter --name is necesary.');
        return;
    }

    person.name = name;

    client.addPerson(person, function (err, replay) {
        console.log(replay.status);
    });
}

function listPeople() {
    client.listPeople({}, function (err, data) {
        if (err) {
            throw err;
        }

        data.people.forEach(function (person) {
            console.log(person.name);
        });
    });
}

function main() {
    var target = argv.target || '127.0.0.1';
    client = new pb.service.myService(target + port, grpc.credentials.createInsecure());

    if (argv.option && argv.option === 'add') {
        addPerson(argv.name);
    } else {
        listPeople();
    }
}

main();