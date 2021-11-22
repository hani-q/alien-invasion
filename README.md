# Alien Invasion

****Requirement**** 

This is my solution to the Tendermint Alien Invasion Technical Challenge, briefly described here:

*"Mad aliens are about to invade the earth and you are tasked with simulating the
invasion.
You are given a map containing the names of cities in the non-existent world of
X. The map is in a file, with one city per line. The city name is first,
followed by 1-4 directions (north, south, east, or west). Each one represents a
road to another city that lies in that direction.
For example:
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
The city and each of the pairs are separated by a single space, and the
directions are separated from their respective cities with an equals (=) sign.
You should create N aliens, where N is specified as a command-line argument.
These aliens start out at random places on the map, and wander around randomly,
following links. Each iteration, the aliens can travel in any of the directions
leading out of a city. In our example above, an alien that starts at Foo can go
north to Bar, west to Baz, or south to Qu-ux.
When two aliens end up in the same place, they fight, and in the process kill
each other and destroy the city. When a city is destroyed, it is removed from
the map, and so are any roads that lead into or out of it.
In our example above, if Bar were destroyed the map would now be something
like:
Foo west=Baz south=Qu-ux
Once a city is destroyed, aliens can no longer travel to or through it. This
may lead to aliens getting "trapped".
You should create a program that reads in the world map, creates N aliens, and
unleashes them. The program should run until all the aliens have been
destroyed, or each alien has moved at least 10,000 times. When two aliens
fight, print out a message like:

Bar has been destroyed by alien 10 and alien 34!
(If you want to give them names, you may, but it is not required.) Once the
program has finished, it should print out whatever is left of the world in the
same format as the input file.
Feel free to make assumptions (for example, that the city names will never
contain numeric characters), but please add comments or assertions describing
the assumptions you are making."*
   


****Lore****

1. Year is 1984
2. X-world is in an alternate dimension (and biologically similar to Earth.)
3. X-world is ruled by an obnoxius Dictator. He for some reason destoryed and recreated all the cities on X World and renamed this world the "X-Grid"
4. The populace in X-world (X-wolrdlings still dont prefer to call it X-Grid) all use flying as means of transportation. Only Cyclists are allowed to use roads.
5. All cities have can have up to 4 roads (in cardinal directions) leading in and out of each city.
6. Alien species is called "Zentradi". 
7. The Zetardi Queen "XX121". Arrived in the Earth Outer orbit and layed her eggs. Which then landed on different cities across X-World
8. The Queen Mother controls the aliens via Telepathic Link
9. Her purpose is to send her children to each city of the world and abosrb the culture
10. After interacting with X-Worlds atmosphere the children can only crawl on Asphalt
11. Zentardi Children have become Feral after interacting with the X-world's atmosphere and wanty to kill each other off. When one Zentradi dies he unleases a nucealear blast that can wipe out an entire city and any other Zentradi
12. X-worldlings dont offer any resitance to the Zentradi invasion
   

****Assumptions****

1. Each city name is unique.
2. City name can be any string literal. "1" is also a valid city name. Who are we to judge X-worldlings.
3. Each city can only have one neghbouring city in one direction. Hence the GRID
4. If "Foo" is to the North of "Baz". Then "Baz" is to the South of "Foo". So adding a link one way.. it will be added both ways between both cities.
5. If "Foo" is destroyed its North-ward road to "Baz" is destroyed. "Baz"'s Southward road to "Foo" is also removed.
6. Inn the start of program each city will have 1 alien and at min have 0.
7. Each city can only handle 2 aliens. these guys are Humongous. 
8. An Alien is HATCHED at the start of the program
9. An Alien is TRAPPED if all roads are destroyed OR all available neighbors are full
10. An Alien is TRAPPED if all roads are destroyed OR all available neighbors are full
11. An Alien is EXHUASTED if number of iterations are over
12. An Alien is DEAD if it encounters another Alien in any other city
13. Alien can go back to the city it just came from.
14. Aliens cannot be greater then the Count of Cities. So in best case sceanrio num Aliens === num Cities.
15. Alien children cannot be less then 2
16. World file will atleast one direction against a mentioned city. This can lead to a problem if One alien has a Road to a city that only has a Road back. It is the Aliens fault for being stupid and getting off-road
17. Each Alien moves with its Hearbeat. One Heartbeat/tick = one Move
18. Each Alien has a different pace with which it Crawls.. Some are young .. some are old Hence the Tick is random
19. Each alien returns a status to Queen mother on her personal channel. 
20. Queen will wait for each Child to return atleast one Status (EXHAUSTED, TRAPPED, DEAD)

***How to Run***
```
go run .\main.go -aliens=10 -numIterations=100  -world_file="./test/world_medium.txt"
```

| flags | description | default |
| ------------- | ------------- | ------------- |
| aliens | Count of the Aliens that need to be spawend. Should be greater then the cities in the map txt file and not less then 2  | 4| 
| numIterations  | How many moves an Individual Alien has  | 10000 |
| world_file  | Sample world map.txt file. Should be the format mentioned in requirements section  | ./test/world_tiny.txt |
| stdout  | Disabled logging to log.txt file in build/ and starts logging to screen | false |

***How to Test***
```
go test -v .\structs\
```