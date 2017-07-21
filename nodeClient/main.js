const grpc = require('grpc');
const argv = require('minimist')(process.argv.slice(2));

const PROTO_PATH = __dirname + '/../service/service.proto';
const pb = grpc.load(PROTO_PATH);

var client = new pb.service.myService('10.4.4.118:50051', grpc.credentials.createInsecure());
var person = {
    name: ''
};

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
    if (argv.option) {
        if (argv.option === 'add') {
            addPerson(argv.name);
        }

        if (argv.option === 'list') {
            listPeople();
        }
    }
}

main();