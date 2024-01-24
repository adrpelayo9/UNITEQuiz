## Live Web Application
A live version of this application can be used at: https://unite-quiz-dhza.vercel.app/
Vercel is used to host the frontend with Next JS, while the backend is hosted by Heroku with Go.

## Pokemon UNITE Quiz Application
This is a project I have been working on in which it tests the knowledge a player may have on the Video Game Pokemon UNITE.
It contains a front-end and a back-end, the front-end using JavaScript with React/Next JS, and the back-end using Golang with Gin. The back-end contains all of the questions which are received by the front-end, a question is taken in and posted to the back-end, where it is evaluated against an answer sheet to check if the answer is correct. If it is, it sends back a response that increases the players score. The score is then evaluated at the end on the front-end to check how well the player did, giving a trophy ranking based on the real in-game rankings within Pokemon UNITE.
I built this application to be mostly dynamic; most if not all of the front-end will dynamically work no matter how large the back-end question and answer slices are. The back-end questions will manually need to be changed for such effect to take place.

## Languages and Frameworks Used
- JavaScript
- Golang
- React
- Next JS
- Gin
- net/http
