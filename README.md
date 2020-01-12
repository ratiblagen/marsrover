# marsrover
InterviewQuestion

For the main.go:
  I did the thing first come to my mind.
  Second rover is on standby until first one finishes its all moves.
  After first one finishes second rover starts to move.
  There is no collison check in this file.
  Only control is out of area control.

For the secondCommit:
  Everything i did for the main stays but i add some features.
   First feature is collison.
  Again second rovers starts its movement after first one finishes but this time its checks last position of the first one and breaks on collison.
  I did not put the control before every movement but put it after every movement.
  So we will know the collison after it happens unfortunately.
  Second feature is movement turns.
  Unlike first thoughts on the subject, i try to make rovers move one by one.
  First one makes a move then second one makes a move, like a chess.
  I checked collison and out of area conditions.
  If rovers instruction is not same length with each other movement stops with smaller instruction.
  
For the lastCommit:
  I try to make my code little bit prettier.
  Tried to get rid of duplicate code blocks.
  And add some comments for the purpose of readability.
  

