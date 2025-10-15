// MongoDB init script to populate the questions collection from quiz-questions.json
// Place this file in localdb/init/ and run with docker-compose up

const fs = require('fs');
const path = '/docker-entrypoint-initdb.d/quiz-questions.json';

let questions = [];
try {
  questions = JSON.parse(fs.readFileSync(path, 'utf8'));
} catch (e) {
  print('Failed to read quiz-questions.json:', e);
}

if (questions.length > 0) {
  db = db.getSiblingDB('sre_quiz');
  db.questions.drop();
  db.questions.insertMany(questions);
  print('Inserted', questions.length, 'questions into sre_quiz.questions');
} else {
  print('No questions loaded.');
}
