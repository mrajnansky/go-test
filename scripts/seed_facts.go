package main

import (
    "fmt"
    "log"
    "github.com/mrajnansky/go-test/database"
    "github.com/mrajnansky/go-test/models"
)

func generateFacts() []models.Fact {
    var facts []models.Fact
    
    geographyFacts := []models.Fact{
        {Question: "What is the capital of France?", Answer: "Paris"},
        {Question: "What is the capital of Germany?", Answer: "Berlin"},
        {Question: "What is the capital of Italy?", Answer: "Rome"},
        {Question: "What is the capital of Spain?", Answer: "Madrid"},
        {Question: "What is the capital of Russia?", Answer: "Moscow"},
        {Question: "What is the capital of China?", Answer: "Beijing"},
        {Question: "What is the capital of Japan?", Answer: "Tokyo"},
        {Question: "What is the capital of India?", Answer: "New Delhi"},
        {Question: "What is the capital of Brazil?", Answer: "Brasília"},
        {Question: "What is the capital of Canada?", Answer: "Ottawa"},
        {Question: "What is the capital of Australia?", Answer: "Canberra"},
        {Question: "What is the capital of Egypt?", Answer: "Cairo"},
        {Question: "What is the capital of South Africa?", Answer: "Cape Town, Pretoria, and Bloemfontein"},
        {Question: "What is the capital of Argentina?", Answer: "Buenos Aires"},
        {Question: "What is the capital of Mexico?", Answer: "Mexico City"},
        {Question: "What is the capital of Turkey?", Answer: "Ankara"},
        {Question: "What is the capital of Thailand?", Answer: "Bangkok"},
        {Question: "What is the capital of Greece?", Answer: "Athens"},
        {Question: "What is the capital of Portugal?", Answer: "Lisbon"},
        {Question: "What is the capital of Norway?", Answer: "Oslo"},
    }
    facts = append(facts, geographyFacts...)
    
    chemistryFacts := []models.Fact{
        {Question: "What is the chemical symbol for gold?", Answer: "Au"},
        {Question: "What is the chemical symbol for silver?", Answer: "Ag"},
        {Question: "What is the chemical symbol for iron?", Answer: "Fe"},
        {Question: "What is the chemical symbol for oxygen?", Answer: "O"},
        {Question: "What is the chemical symbol for hydrogen?", Answer: "H"},
        {Question: "What is the chemical symbol for carbon?", Answer: "C"},
        {Question: "What is the chemical symbol for nitrogen?", Answer: "N"},
        {Question: "What is the chemical symbol for helium?", Answer: "He"},
        {Question: "What is the chemical symbol for sodium?", Answer: "Na"},
        {Question: "What is the chemical symbol for chlorine?", Answer: "Cl"},
        {Question: "What is the chemical formula for water?", Answer: "H2O"},
        {Question: "What is the chemical formula for carbon dioxide?", Answer: "CO2"},
        {Question: "What is the chemical formula for methane?", Answer: "CH4"},
        {Question: "What is the chemical formula for ammonia?", Answer: "NH3"},
        {Question: "What is the chemical formula for table salt?", Answer: "NaCl"},
        {Question: "What is the atomic number of hydrogen?", Answer: "1"},
        {Question: "What is the atomic number of carbon?", Answer: "6"},
        {Question: "What is the atomic number of oxygen?", Answer: "8"},
        {Question: "What is the atomic number of gold?", Answer: "79"},
        {Question: "What is the atomic number of uranium?", Answer: "92"},
    }
    facts = append(facts, chemistryFacts...)
    
    physicsFacts := []models.Fact{
        {Question: "What is the speed of light?", Answer: "299,792,458 meters per second"},
        {Question: "What is the acceleration due to gravity on Earth?", Answer: "9.8 m/s²"},
        {Question: "What is the boiling point of water?", Answer: "100 degrees Celsius"},
        {Question: "What is the freezing point of water?", Answer: "0 degrees Celsius"},
        {Question: "What is the melting point of ice?", Answer: "0 degrees Celsius"},
        {Question: "What is absolute zero?", Answer: "-273.15 degrees Celsius"},
        {Question: "What is the speed of sound in air?", Answer: "343 meters per second"},
        {Question: "What is Newton's first law of motion?", Answer: "An object at rest stays at rest unless acted upon by a force"},
        {Question: "What is Newton's second law of motion?", Answer: "Force equals mass times acceleration (F=ma)"},
        {Question: "What is Newton's third law of motion?", Answer: "For every action, there is an equal and opposite reaction"},
        {Question: "What is the unit of force?", Answer: "Newton"},
        {Question: "What is the unit of energy?", Answer: "Joule"},
        {Question: "What is the unit of power?", Answer: "Watt"},
        {Question: "What is the unit of electric current?", Answer: "Ampere"},
        {Question: "What is the unit of electric potential?", Answer: "Volt"},
        {Question: "What is Ohm's law?", Answer: "V = I × R"},
        {Question: "What is the electromagnetic spectrum?", Answer: "The range of all electromagnetic radiation"},
        {Question: "What is a photon?", Answer: "A particle of light"},
        {Question: "What is Einstein's famous equation?", Answer: "E=mc²"},
        {Question: "What is the Planck constant?", Answer: "6.626 × 10⁻³⁴ J·s"},
    }
    facts = append(facts, physicsFacts...)
    
    biologyFacts := []models.Fact{
        {Question: "How many bones are in the human body?", Answer: "206"},
        {Question: "How many chambers does a human heart have?", Answer: "Four"},
        {Question: "How many teeth does an adult human have?", Answer: "32"},
        {Question: "How many ribs does a human have?", Answer: "24"},
        {Question: "How many vertebrae are in the human spine?", Answer: "33"},
        {Question: "What is the largest organ in the human body?", Answer: "Skin"},
        {Question: "What is the smallest muscle in the human body?", Answer: "Stapedius"},
        {Question: "What is the largest bone in the human body?", Answer: "Femur"},
        {Question: "What is the longest bone in the human body?", Answer: "Femur"},
        {Question: "What is the smallest bone in the human body?", Answer: "Stapes"},
        {Question: "What gas do plants absorb from the atmosphere?", Answer: "Carbon dioxide"},
        {Question: "What gas do plants release during photosynthesis?", Answer: "Oxygen"},
        {Question: "What is the process by which plants make food?", Answer: "Photosynthesis"},
        {Question: "What is DNA?", Answer: "Deoxyribonucleic acid"},
        {Question: "What is RNA?", Answer: "Ribonucleic acid"},
        {Question: "How many chromosomes do humans have?", Answer: "46"},
        {Question: "What is the powerhouse of the cell?", Answer: "Mitochondria"},
        {Question: "What is the control center of the cell?", Answer: "Nucleus"},
        {Question: "What is the basic unit of life?", Answer: "Cell"},
        {Question: "What is evolution?", Answer: "The process of change in species over time"},
    }
    facts = append(facts, biologyFacts...)
    
    historyFacts := []models.Fact{
        {Question: "What year did World War II end?", Answer: "1945"},
        {Question: "What year was the first moon landing?", Answer: "1969"},
        {Question: "What year did the Berlin Wall fall?", Answer: "1989"},
        {Question: "What year did the Titanic sink?", Answer: "1912"},
        {Question: "What year did World War I begin?", Answer: "1914"},
        {Question: "What year did the American Civil War end?", Answer: "1865"},
        {Question: "What year was the Declaration of Independence signed?", Answer: "1776"},
        {Question: "What year did Christopher Columbus reach the Americas?", Answer: "1492"},
        {Question: "What year did the French Revolution begin?", Answer: "1789"},
        {Question: "What year did the Russian Revolution occur?", Answer: "1917"},
        {Question: "Who was the first President of the United States?", Answer: "George Washington"},
        {Question: "Who was the 16th President of the United States?", Answer: "Abraham Lincoln"},
        {Question: "Who was the longest-serving President of the United States?", Answer: "Franklin D. Roosevelt"},
        {Question: "Who discovered America?", Answer: "Christopher Columbus"},
        {Question: "Who discovered gravity?", Answer: "Isaac Newton"},
        {Question: "Who discovered electricity?", Answer: "Benjamin Franklin"},
        {Question: "Who discovered penicillin?", Answer: "Alexander Fleming"},
        {Question: "Who invented the telephone?", Answer: "Alexander Graham Bell"},
        {Question: "Who invented the light bulb?", Answer: "Thomas Edison"},
        {Question: "Who invented the airplane?", Answer: "Wright Brothers"},
    }
    facts = append(facts, historyFacts...)
    
    literatureFacts := []models.Fact{
        {Question: "Who wrote Romeo and Juliet?", Answer: "William Shakespeare"},
        {Question: "Who wrote Pride and Prejudice?", Answer: "Jane Austen"},
        {Question: "Who wrote 1984?", Answer: "George Orwell"},
        {Question: "Who wrote To Kill a Mockingbird?", Answer: "Harper Lee"},
        {Question: "Who wrote The Great Gatsby?", Answer: "F. Scott Fitzgerald"},
        {Question: "Who wrote Harry Potter?", Answer: "J.K. Rowling"},
        {Question: "Who wrote Lord of the Rings?", Answer: "J.R.R. Tolkien"},
        {Question: "Who wrote Moby Dick?", Answer: "Herman Melville"},
        {Question: "Who wrote The Catcher in the Rye?", Answer: "J.D. Salinger"},
        {Question: "Who wrote War and Peace?", Answer: "Leo Tolstoy"},
        {Question: "Who wrote Crime and Punishment?", Answer: "Fyodor Dostoevsky"},
        {Question: "Who wrote One Hundred Years of Solitude?", Answer: "Gabriel García Márquez"},
        {Question: "Who wrote The Odyssey?", Answer: "Homer"},
        {Question: "Who wrote The Iliad?", Answer: "Homer"},
        {Question: "Who wrote Don Quixote?", Answer: "Miguel de Cervantes"},
        {Question: "Who wrote Hamlet?", Answer: "William Shakespeare"},
        {Question: "Who wrote Macbeth?", Answer: "William Shakespeare"},
        {Question: "Who wrote Othello?", Answer: "William Shakespeare"},
        {Question: "Who wrote King Lear?", Answer: "William Shakespeare"},
        {Question: "Who wrote The Tempest?", Answer: "William Shakespeare"},
    }
    facts = append(facts, literatureFacts...)
    
    artFacts := []models.Fact{
        {Question: "Who painted the Mona Lisa?", Answer: "Leonardo da Vinci"},
        {Question: "Who painted Starry Night?", Answer: "Vincent van Gogh"},
        {Question: "Who painted The Last Supper?", Answer: "Leonardo da Vinci"},
        {Question: "Who painted The Scream?", Answer: "Edvard Munch"},
        {Question: "Who painted Guernica?", Answer: "Pablo Picasso"},
        {Question: "Who painted The Persistence of Memory?", Answer: "Salvador Dalí"},
        {Question: "Who painted The Girl with a Pearl Earring?", Answer: "Johannes Vermeer"},
        {Question: "Who painted The Birth of Venus?", Answer: "Sandro Botticelli"},
        {Question: "Who painted American Gothic?", Answer: "Grant Wood"},
        {Question: "Who painted The Great Wave off Kanagawa?", Answer: "Hokusai"},
        {Question: "Who sculpted David?", Answer: "Michelangelo"},
        {Question: "Who sculpted The Thinker?", Answer: "Auguste Rodin"},
        {Question: "Who painted the ceiling of the Sistine Chapel?", Answer: "Michelangelo"},
        {Question: "Who painted Water Lilies?", Answer: "Claude Monet"},
        {Question: "Who painted Sunflowers?", Answer: "Vincent van Gogh"},
        {Question: "Who painted The Starry Night Over the Rhône?", Answer: "Vincent van Gogh"},
        {Question: "Who painted Las Meninas?", Answer: "Diego Velázquez"},
        {Question: "Who painted The Night Watch?", Answer: "Rembrandt"},
        {Question: "Who painted Liberty Leading the People?", Answer: "Eugène Delacroix"},
        {Question: "Who painted A Sunday on La Grande Jatte?", Answer: "Georges Seurat"},
    }
    facts = append(facts, artFacts...)
    
    musicFacts := []models.Fact{
        {Question: "Who composed The Four Seasons?", Answer: "Antonio Vivaldi"},
        {Question: "Who composed Beethoven's 9th Symphony?", Answer: "Ludwig van Beethoven"},
        {Question: "Who composed The Magic Flute?", Answer: "Wolfgang Amadeus Mozart"},
        {Question: "Who composed Carmen?", Answer: "Georges Bizet"},
        {Question: "Who composed The Nutcracker?", Answer: "Pyotr Ilyich Tchaikovsky"},
        {Question: "Who composed Swan Lake?", Answer: "Pyotr Ilyich Tchaikovsky"},
        {Question: "Who composed The Ring of the Nibelung?", Answer: "Richard Wagner"},
        {Question: "Who composed The Barber of Seville?", Answer: "Gioachino Rossini"},
        {Question: "Who composed Don Giovanni?", Answer: "Wolfgang Amadeus Mozart"},
        {Question: "Who composed Requiem?", Answer: "Wolfgang Amadeus Mozart"},
        {Question: "How many strings does a violin have?", Answer: "Four"},
        {Question: "How many strings does a guitar have?", Answer: "Six"},
        {Question: "How many keys are on a standard piano?", Answer: "88"},
        {Question: "What is the highest female singing voice?", Answer: "Soprano"},
        {Question: "What is the lowest male singing voice?", Answer: "Bass"},
        {Question: "What is a group of four musicians called?", Answer: "Quartet"},
        {Question: "What is a group of three musicians called?", Answer: "Trio"},
        {Question: "What is the term for the speed of music?", Answer: "Tempo"},
        {Question: "What is the term for the loudness of music?", Answer: "Dynamics"},
        {Question: "What is the term for the highness or lowness of sound?", Answer: "Pitch"},
    }
    facts = append(facts, musicFacts...)
    
    astronomyFacts := []models.Fact{
        {Question: "What is the largest planet in our solar system?", Answer: "Jupiter"},
        {Question: "What is the smallest planet in our solar system?", Answer: "Mercury"},
        {Question: "What is the second largest planet in our solar system?", Answer: "Saturn"},
        {Question: "What planet is known as the Red Planet?", Answer: "Mars"},
        {Question: "What planet is closest to the Sun?", Answer: "Mercury"},
        {Question: "What is the third planet from the Sun?", Answer: "Earth"},
        {Question: "What is the hottest planet in our solar system?", Answer: "Venus"},
        {Question: "What planet has the most moons?", Answer: "Jupiter"},
        {Question: "How many moons does Earth have?", Answer: "One"},
        {Question: "What is the name of Earth's moon?", Answer: "The Moon"},
        {Question: "What is the name of Mars' largest moon?", Answer: "Phobos"},
        {Question: "What is the largest moon in our solar system?", Answer: "Ganymede"},
        {Question: "What is the closest star to Earth?", Answer: "The Sun"},
        {Question: "What is the second closest star to Earth?", Answer: "Proxima Centauri"},
        {Question: "What galaxy do we live in?", Answer: "The Milky Way"},
        {Question: "What is the nearest galaxy to the Milky Way?", Answer: "Andromeda"},
        {Question: "What is a black hole?", Answer: "A region of space with gravity so strong that nothing can escape"},
        {Question: "What is a supernova?", Answer: "The explosive death of a massive star"},
        {Question: "What is a nebula?", Answer: "A cloud of gas and dust in space"},
        {Question: "What is the Big Bang?", Answer: "The theory of how the universe began"},
    }
    facts = append(facts, astronomyFacts...)
    
    mathFacts := []models.Fact{
        {Question: "What is pi approximately equal to?", Answer: "3.14159"},
        {Question: "What is the square root of 144?", Answer: "12"},
        {Question: "What is 2 to the power of 10?", Answer: "1024"},
        {Question: "What is the sum of angles in a triangle?", Answer: "180 degrees"},
        {Question: "What is the sum of angles in a quadrilateral?", Answer: "360 degrees"},
        {Question: "How many sides does a hexagon have?", Answer: "Six"},
        {Question: "How many sides does a pentagon have?", Answer: "Five"},
        {Question: "How many sides does an octagon have?", Answer: "Eight"},
        {Question: "What is the formula for the area of a circle?", Answer: "πr²"},
        {Question: "What is the formula for the circumference of a circle?", Answer: "2πr"},
        {Question: "What is the Pythagorean theorem?", Answer: "a² + b² = c²"},
        {Question: "What is the golden ratio approximately?", Answer: "1.618"},
        {Question: "What is zero factorial?", Answer: "1"},
        {Question: "What is the derivative of x²?", Answer: "2x"},
        {Question: "What is the integral of 1/x?", Answer: "ln(x) + C"},
        {Question: "What is Euler's number approximately?", Answer: "2.718"},
        {Question: "What is the smallest prime number?", Answer: "2"},
        {Question: "What is the largest prime number less than 100?", Answer: "97"},
        {Question: "What is a perfect number?", Answer: "A number equal to the sum of its proper divisors"},
        {Question: "What is the first perfect number?", Answer: "6"},
    }
    facts = append(facts, mathFacts...)
    
    
    animalFacts := []models.Fact{
        {Question: "What is the largest mammal in the world?", Answer: "Blue whale"},
        {Question: "What is the fastest land animal?", Answer: "Cheetah"},
        {Question: "What is the slowest mammal?", Answer: "Sloth"},
        {Question: "What is the largest bird in the world?", Answer: "Ostrich"},
        {Question: "What is the smallest bird in the world?", Answer: "Bee hummingbird"},
        {Question: "What is the largest fish in the world?", Answer: "Whale shark"},
        {Question: "What is the fastest marine animal?", Answer: "Sailfish"},
        {Question: "What is the smallest dog breed?", Answer: "Chihuahua"},
        {Question: "What is the largest dog breed?", Answer: "Great Dane"},
        {Question: "How many legs does a spider have?", Answer: "Eight"},
        {Question: "How many eyes does a spider have?", Answer: "Eight"},
        {Question: "How many tentacles does an octopus have?", Answer: "Eight"},
        {Question: "How many wings does a bee have?", Answer: "Four"},
        {Question: "How many chambers does a fish heart have?", Answer: "Two"},
        {Question: "What is the hardest natural substance on Earth?", Answer: "Diamond"},
        {Question: "What is the softest mineral?", Answer: "Talc"},
        {Question: "What is the largest desert in the world?", Answer: "Antarctica"},
        {Question: "What is the highest waterfall in the world?", Answer: "Angel Falls"},
        {Question: "What is the deepest ocean trench?", Answer: "Mariana Trench"},
        {Question: "What is the longest river in the world?", Answer: "The Nile River"},
    }
    facts = append(facts, animalFacts...)
    
    technologyFacts := []models.Fact{
        {Question: "Who invented the World Wide Web?", Answer: "Tim Berners-Lee"},
        {Question: "Who invented the printing press?", Answer: "Johannes Gutenberg"},
        {Question: "Who invented the steam engine?", Answer: "James Watt"},
        {Question: "Who invented the radio?", Answer: "Guglielmo Marconi"},
        {Question: "Who invented the television?", Answer: "John Logie Baird"},
        {Question: "Who founded Microsoft?", Answer: "Bill Gates and Paul Allen"},
        {Question: "Who founded Apple?", Answer: "Steve Jobs, Steve Wozniak, and Ronald Wayne"},
        {Question: "Who founded Google?", Answer: "Larry Page and Sergey Brin"},
        {Question: "Who founded Facebook?", Answer: "Mark Zuckerberg"},
        {Question: "Who founded Amazon?", Answer: "Jeff Bezos"},
        {Question: "What does CPU stand for?", Answer: "Central Processing Unit"},
        {Question: "What does RAM stand for?", Answer: "Random Access Memory"},
        {Question: "What does GPU stand for?", Answer: "Graphics Processing Unit"},
        {Question: "What does HTML stand for?", Answer: "HyperText Markup Language"},
        {Question: "What does HTTP stand for?", Answer: "HyperText Transfer Protocol"},
        {Question: "What does URL stand for?", Answer: "Uniform Resource Locator"},
        {Question: "What does AI stand for?", Answer: "Artificial Intelligence"},
        {Question: "What does VR stand for?", Answer: "Virtual Reality"},
        {Question: "What does AR stand for?", Answer: "Augmented Reality"},
        {Question: "What does IoT stand for?", Answer: "Internet of Things"},
    }
    facts = append(facts, technologyFacts...)
    
    foodFacts := []models.Fact{
        {Question: "What is the main ingredient in guacamole?", Answer: "Avocado"},
        {Question: "What is the main ingredient in hummus?", Answer: "Chickpeas"},
        {Question: "What is the hottest chili pepper?", Answer: "Carolina Reaper"},
        {Question: "What is the most expensive spice by weight?", Answer: "Saffron"},
        {Question: "What is the main ingredient in traditional Japanese miso soup?", Answer: "Miso paste"},
        {Question: "What is the primary ingredient in traditional Greek tzatziki?", Answer: "Yogurt"},
        {Question: "What is the main ingredient in pesto?", Answer: "Basil"},
        {Question: "What is the fermented cabbage dish from Korea called?", Answer: "Kimchi"},
        {Question: "What is the Italian rice dish called?", Answer: "Risotto"},
        {Question: "What is the Spanish rice dish called?", Answer: "Paella"},
        {Question: "What is the French term for cooking vegetables until soft?", Answer: "Sauté"},
        {Question: "What is the cooking method where food is cooked in its own fat?", Answer: "Confit"},
        {Question: "What is the cooking method involving dry heat in an oven?", Answer: "Roasting"},
        {Question: "What is the cooking method involving moist heat with liquid?", Answer: "Braising"},
        {Question: "What is the Maillard reaction?", Answer: "The browning reaction between amino acids and sugars"},
        {Question: "What temperature should poultry be cooked to for safety?", Answer: "165°F (74°C)"},
        {Question: "What is the difference between baking soda and baking powder?", Answer: "Baking soda needs acid to activate; baking powder has acid built in"},
        {Question: "What is gluten?", Answer: "A protein found in wheat, barley, and rye"},
        {Question: "What is umami?", Answer: "The fifth taste, often described as savory"},
        {Question: "What is the smoking point of olive oil?", Answer: "About 375°F (190°C)"},
    }
    facts = append(facts, foodFacts...)
    
    sportsFacts := []models.Fact{
        {Question: "How many players are on a basketball team on the court?", Answer: "Five"},
        {Question: "How many players are on a soccer team on the field?", Answer: "Eleven"},
        {Question: "How many players are on a baseball team on the field?", Answer: "Nine"},
        {Question: "How many periods are in a hockey game?", Answer: "Three"},
        {Question: "How many quarters are in a football game?", Answer: "Four"},
        {Question: "What is the maximum score in ten-pin bowling?", Answer: "300"},
        {Question: "How many holes are on a standard golf course?", Answer: "18"},
        {Question: "What is the distance of a marathon?", Answer: "26.2 miles or 42.195 kilometers"},
        {Question: "How many rings are on the Olympic flag?", Answer: "Five"},
        {Question: "What are the five Olympic ring colors?", Answer: "Blue, yellow, black, green, and red"},
        {Question: "How often are the Summer Olympics held?", Answer: "Every four years"},
        {Question: "How often are the Winter Olympics held?", Answer: "Every four years"},
        {Question: "What is the highest possible hand in poker?", Answer: "Royal flush"},
        {Question: "How many cards are in a standard deck?", Answer: "52"},
        {Question: "How many squares are on a chessboard?", Answer: "64"},
        {Question: "How many pieces does each player start with in chess?", Answer: "16"},
        {Question: "What is the object ball you must sink last in pool?", Answer: "The 8-ball"},
        {Question: "How many pins are there in bowling?", Answer: "10"},
        {Question: "What is par for most golf holes?", Answer: "3, 4, or 5"},
        {Question: "How many points is a touchdown worth in American football?", Answer: "Six"},
    }
    facts = append(facts, sportsFacts...)
    
    
    recordFacts := []models.Fact{
        {Question: "What is the tallest mountain in the world?", Answer: "Mount Everest"},
        {Question: "What is the tallest building in the world?", Answer: "Burj Khalifa"},
        {Question: "What is the longest wall in the world?", Answer: "Great Wall of China"},
        {Question: "What is the largest ocean on Earth?", Answer: "Pacific Ocean"},
        {Question: "What is the smallest ocean on Earth?", Answer: "Arctic Ocean"},
        {Question: "What is the largest lake in the world?", Answer: "Caspian Sea"},
        {Question: "What is the deepest lake in the world?", Answer: "Lake Baikal"},
        {Question: "What is the largest island in the world?", Answer: "Greenland"},
        {Question: "What is the smallest country in the world?", Answer: "Vatican City"},
        {Question: "What is the largest country in the world?", Answer: "Russia"},
        {Question: "What is the most populous country in the world?", Answer: "China"},
        {Question: "What is the most populous city in the world?", Answer: "Tokyo"},
        {Question: "What is the driest place on Earth?", Answer: "Atacama Desert"},
        {Question: "What is the wettest place on Earth?", Answer: "Mawsynram, India"},
        {Question: "What is the hottest place on Earth?", Answer: "Death Valley, California"},
        {Question: "What is the coldest place on Earth?", Answer: "Antarctica"},
        {Question: "What is the oldest living tree?", Answer: "Methuselah (a bristlecone pine)"},
        {Question: "What is the largest living thing on Earth?", Answer: "Armillaria ostoyae (honey fungus)"},
        {Question: "What is the fastest bird in flight?", Answer: "Peregrine falcon"},
        {Question: "What is the largest spider in the world?", Answer: "Goliath birdeater"},
    }
    facts = append(facts, recordFacts...)
    
    timeFacts := []models.Fact{
        {Question: "How many seconds are in a minute?", Answer: "60"},
        {Question: "How many minutes are in an hour?", Answer: "60"},
        {Question: "How many hours are in a day?", Answer: "24"},
        {Question: "How many days are in a week?", Answer: "Seven"},
        {Question: "How many days are in a year?", Answer: "365"},
        {Question: "How many days are in a leap year?", Answer: "366"},
        {Question: "How many seconds are in an hour?", Answer: "3,600"},
        {Question: "How many minutes are in a full day?", Answer: "1,440"},
        {Question: "How many months have 31 days?", Answer: "Seven"},
        {Question: "How many months have 30 days?", Answer: "Four"},
        {Question: "Which month has the fewest days?", Answer: "February"},
        {Question: "How many weeks are in a year?", Answer: "52"},
        {Question: "What is the first month of the year?", Answer: "January"},
        {Question: "What is the last month of the year?", Answer: "December"},
        {Question: "What is the shortest day of the year called?", Answer: "Winter solstice"},
        {Question: "What is the longest day of the year called?", Answer: "Summer solstice"},
        {Question: "How often does a leap year occur?", Answer: "Every four years"},
        {Question: "What determines a leap year?", Answer: "Years divisible by 4, except century years not divisible by 400"},
        {Question: "What is GMT?", Answer: "Greenwich Mean Time"},
        {Question: "How many time zones are there in the world?", Answer: "24"},
    }
    facts = append(facts, timeFacts...)
    
    currencyFacts := []models.Fact{
        {Question: "What is the currency of Japan?", Answer: "Yen"},
        {Question: "What is the currency of the United Kingdom?", Answer: "Pound Sterling"},
        {Question: "What is the currency of China?", Answer: "Yuan"},
        {Question: "What is the currency of Mexico?", Answer: "Peso"},
        {Question: "What is the currency of India?", Answer: "Rupee"},
        {Question: "What is the currency of Russia?", Answer: "Ruble"},
        {Question: "What is the currency of South Korea?", Answer: "Won"},
        {Question: "What is the currency of Switzerland?", Answer: "Swiss Franc"},
        {Question: "What is the currency of Canada?", Answer: "Canadian Dollar"},
        {Question: "What is the currency of Australia?", Answer: "Australian Dollar"},
        {Question: "What is the currency of Brazil?", Answer: "Real"},
        {Question: "What is the currency of South Africa?", Answer: "Rand"},
        {Question: "What is the currency of Turkey?", Answer: "Lira"},
        {Question: "What is the currency of Norway?", Answer: "Krone"},
        {Question: "What is the currency of Sweden?", Answer: "Krona"},
        {Question: "What does GDP stand for?", Answer: "Gross Domestic Product"},
        {Question: "What does CEO stand for?", Answer: "Chief Executive Officer"},
        {Question: "What does IPO stand for?", Answer: "Initial Public Offering"},
        {Question: "What does NYSE stand for?", Answer: "New York Stock Exchange"},
        {Question: "What does NASDAQ stand for?", Answer: "National Association of Securities Dealers Automated Quotations"},
    }
    facts = append(facts, currencyFacts...)
    
    bodyFacts := []models.Fact{
        {Question: "What is the largest organ in the human body?", Answer: "Skin"},
        {Question: "What is the smallest organ in the human body?", Answer: "Pineal gland"},
        {Question: "What is the hardest substance in the human body?", Answer: "Tooth enamel"},
        {Question: "How many muscles are in the human body?", Answer: "About 600"},
        {Question: "What is the strongest muscle in the human body?", Answer: "Masseter (jaw muscle)"},
        {Question: "What is the longest nerve in the human body?", Answer: "Sciatic nerve"},
        {Question: "How many taste buds does the average person have?", Answer: "About 10,000"},
        {Question: "How many times does the heart beat per day?", Answer: "About 100,000"},
        {Question: "How much blood does the human body contain?", Answer: "About 5 liters"},
        {Question: "How many red blood cells are in the human body?", Answer: "About 25 trillion"},
        {Question: "What is the lifespan of a red blood cell?", Answer: "About 120 days"},
        {Question: "How many breaths does a person take per day?", Answer: "About 20,000"},
        {Question: "What percentage of the human body is water?", Answer: "About 60%"},
        {Question: "How many hairs are on the average human head?", Answer: "About 100,000"},
        {Question: "What is the fastest healing part of the body?", Answer: "Tongue"},
        {Question: "What is the slowest healing part of the body?", Answer: "Cartilage"},
        {Question: "How many sweat glands does the human body have?", Answer: "About 2-4 million"},
        {Question: "What is the normal human body temperature?", Answer: "98.6°F or 37°C"},
        {Question: "How many pairs of cranial nerves are there?", Answer: "12"},
        {Question: "What is the medical term for the kneecap?", Answer: "Patella"},
    }
    facts = append(facts, bodyFacts...)
    
    
    moreFacts := generateAdditionalFacts(5000 - len(facts))
    facts = append(facts, moreFacts...)
    
    return facts
}

func generateAdditionalFacts(needed int) []models.Fact {
    var additionalFacts []models.Fact
    
    for i := 1; i <= needed/10 && len(additionalFacts) < needed; i++ {
        additionalFacts = append(additionalFacts, models.Fact{
            Question: fmt.Sprintf("What is %d multiplied by %d?", i, i),
            Answer:   fmt.Sprintf("%d", i*i),
        })
    }
    
    countries := []struct{country, capital string}{
        {"Afghanistan", "Kabul"}, {"Albania", "Tirana"}, {"Algeria", "Algiers"},
        {"Angola", "Luanda"}, {"Armenia", "Yerevan"}, {"Austria", "Vienna"},
        {"Azerbaijan", "Baku"}, {"Bahrain", "Manama"}, {"Bangladesh", "Dhaka"},
        {"Belarus", "Minsk"}, {"Belgium", "Brussels"}, {"Bolivia", "La Paz"},
        {"Bosnia and Herzegovina", "Sarajevo"}, {"Bulgaria", "Sofia"}, {"Cambodia", "Phnom Penh"},
        {"Chile", "Santiago"}, {"Colombia", "Bogotá"}, {"Croatia", "Zagreb"},
        {"Czech Republic", "Prague"}, {"Denmark", "Copenhagen"}, {"Ecuador", "Quito"},
        {"Estonia", "Tallinn"}, {"Ethiopia", "Addis Ababa"}, {"Finland", "Helsinki"},
        {"Georgia", "Tbilisi"}, {"Ghana", "Accra"}, {"Hungary", "Budapest"},
        {"Iceland", "Reykjavik"}, {"Indonesia", "Jakarta"}, {"Iran", "Tehran"},
        {"Iraq", "Baghdad"}, {"Ireland", "Dublin"}, {"Israel", "Jerusalem"},
        {"Jordan", "Amman"}, {"Kazakhstan", "Nur-Sultan"}, {"Kenya", "Nairobi"},
        {"Latvia", "Riga"}, {"Lebanon", "Beirut"}, {"Lithuania", "Vilnius"},
        {"Luxembourg", "Luxembourg City"}, {"Malaysia", "Kuala Lumpur"}, {"Morocco", "Rabat"},
        {"Nepal", "Kathmandu"}, {"Netherlands", "Amsterdam"}, {"New Zealand", "Wellington"},
        {"Nigeria", "Abuja"}, {"Pakistan", "Islamabad"}, {"Peru", "Lima"},
        {"Philippines", "Manila"}, {"Poland", "Warsaw"}, {"Romania", "Bucharest"},
        {"Saudi Arabia", "Riyadh"}, {"Serbia", "Belgrade"}, {"Singapore", "Singapore"},
        {"Slovakia", "Bratislava"}, {"Slovenia", "Ljubljana"}, {"Sri Lanka", "Colombo"},
        {"Sudan", "Khartoum"}, {"Switzerland", "Bern"}, {"Syria", "Damascus"},
        {"Tunisia", "Tunis"}, {"Ukraine", "Kyiv"}, {"United Arab Emirates", "Abu Dhabi"},
        {"Uruguay", "Montevideo"}, {"Venezuela", "Caracas"}, {"Vietnam", "Hanoi"},
        {"Yemen", "Sana'a"}, {"Zambia", "Lusaka"}, {"Zimbabwe", "Harare"},
    }
    
    for _, country := range countries {
        if len(additionalFacts) >= needed { break }
        additionalFacts = append(additionalFacts, models.Fact{
            Question: fmt.Sprintf("What is the capital of %s?", country.country),
            Answer:   country.capital,
        })
    }
    
    elements := []struct{element, symbol string}{
        {"Lithium", "Li"}, {"Beryllium", "Be"}, {"Boron", "B"}, {"Fluorine", "F"},
        {"Neon", "Ne"}, {"Magnesium", "Mg"}, {"Aluminum", "Al"}, {"Silicon", "Si"},
        {"Phosphorus", "P"}, {"Sulfur", "S"}, {"Argon", "Ar"}, {"Potassium", "K"},
        {"Calcium", "Ca"}, {"Scandium", "Sc"}, {"Titanium", "Ti"}, {"Vanadium", "V"},
        {"Chromium", "Cr"}, {"Manganese", "Mn"}, {"Nickel", "Ni"}, {"Copper", "Cu"},
        {"Zinc", "Zn"}, {"Gallium", "Ga"}, {"Germanium", "Ge"}, {"Arsenic", "As"},
        {"Selenium", "Se"}, {"Bromine", "Br"}, {"Krypton", "Kr"}, {"Rubidium", "Rb"},
        {"Strontium", "Sr"}, {"Yttrium", "Y"}, {"Zirconium", "Zr"}, {"Niobium", "Nb"},
        {"Molybdenum", "Mo"}, {"Technetium", "Tc"}, {"Ruthenium", "Ru"}, {"Rhodium", "Rh"},
        {"Palladium", "Pd"}, {"Cadmium", "Cd"}, {"Indium", "In"}, {"Tin", "Sn"},
        {"Antimony", "Sb"}, {"Tellurium", "Te"}, {"Iodine", "I"}, {"Xenon", "Xe"},
        {"Cesium", "Cs"}, {"Barium", "Ba"}, {"Lanthanum", "La"}, {"Cerium", "Ce"},
        {"Praseodymium", "Pr"}, {"Neodymium", "Nd"}, {"Promethium", "Pm"}, {"Samarium", "Sm"},
        {"Europium", "Eu"}, {"Gadolinium", "Gd"}, {"Terbium", "Tb"}, {"Dysprosium", "Dy"},
        {"Holmium", "Ho"}, {"Erbium", "Er"}, {"Thulium", "Tm"}, {"Ytterbium", "Yb"},
        {"Lutetium", "Lu"}, {"Hafnium", "Hf"}, {"Tantalum", "Ta"}, {"Tungsten", "W"},
        {"Rhenium", "Re"}, {"Osmium", "Os"}, {"Iridium", "Ir"}, {"Platinum", "Pt"},
        {"Mercury", "Hg"}, {"Thallium", "Tl"}, {"Lead", "Pb"}, {"Bismuth", "Bi"},
        {"Polonium", "Po"}, {"Astatine", "At"}, {"Radon", "Rn"}, {"Francium", "Fr"},
        {"Radium", "Ra"}, {"Actinium", "Ac"}, {"Thorium", "Th"}, {"Protactinium", "Pa"},
    }
    
    for _, element := range elements {
        if len(additionalFacts) >= needed { break }
        additionalFacts = append(additionalFacts, models.Fact{
            Question: fmt.Sprintf("What is the chemical symbol for %s?", element.element),
            Answer:   element.symbol,
        })
    }
    
    animals := []struct{animal, fact string}{
        {"Elephant", "Can weigh up to 14,000 pounds"},
        {"Giraffe", "Can be up to 18 feet tall"},
        {"Hummingbird", "Can flap wings up to 80 times per second"},
        {"Koala", "Sleeps up to 22 hours per day"},
        {"Cheetah", "Can run up to 70 mph"},
        {"Penguin", "Cannot fly but are excellent swimmers"},
        {"Kangaroo", "Cannot walk backwards"},
        {"Flamingo", "Gets pink color from eating shrimp"},
        {"Owl", "Can rotate head 270 degrees"},
        {"Dolphin", "Uses echolocation to navigate"},
        {"Shark", "Has been around for 400 million years"},
        {"Turtle", "Can live over 100 years"},
        {"Butterfly", "Tastes with its feet"},
        {"Bee", "Visits up to 5,000 flowers in one day"},
        {"Ant", "Can lift 50 times its body weight"},
        {"Spider", "All spiders are carnivorous"},
        {"Frog", "Breathes through its skin"},
        {"Snake", "Cannot close its eyes"},
        {"Crocodile", "Has remained unchanged for 200 million years"},
        {"Bat", "Is the only mammal that can fly"},
    }
    
    for _, animal := range animals {
        if len(additionalFacts) >= needed { break }
        additionalFacts = append(additionalFacts, models.Fact{
            Question: fmt.Sprintf("What is a fact about %ss?", animal.animal),
            Answer:   animal.fact,
        })
    }
    
    for i := 1; i <= 100 && len(additionalFacts) < needed; i++ {
        additionalFacts = append(additionalFacts, models.Fact{
            Question: fmt.Sprintf("What is the square of %d?", i),
            Answer:   fmt.Sprintf("%d", i*i),
        })
    }
    
    planetFacts := []models.Fact{
        {Question: "How many moons does Mercury have?", Answer: "Zero"},
        {Question: "How many moons does Venus have?", Answer: "Zero"},
        {Question: "How many moons does Mars have?", Answer: "Two"},
        {Question: "How many moons does Jupiter have?", Answer: "Over 80"},
        {Question: "How many moons does Saturn have?", Answer: "Over 80"},
        {Question: "How many moons does Uranus have?", Answer: "27"},
        {Question: "How many moons does Neptune have?", Answer: "14"},
        {Question: "What is the diameter of Earth?", Answer: "12,742 kilometers"},
        {Question: "What is the diameter of the Moon?", Answer: "3,474 kilometers"},
        {Question: "How long is a day on Venus?", Answer: "243 Earth days"},
        {Question: "How long is a year on Mercury?", Answer: "88 Earth days"},
        {Question: "How long is a year on Mars?", Answer: "687 Earth days"},
        {Question: "What is the Great Red Spot?", Answer: "A giant storm on Jupiter"},
        {Question: "What are Saturn's rings made of?", Answer: "Ice and rock particles"},
        {Question: "What is unique about Uranus's rotation?", Answer: "It rotates on its side"},
        {Question: "What is the windiest planet?", Answer: "Neptune"},
        {Question: "What is the asteroid belt?", Answer: "A region of space between Mars and Jupiter"},
        {Question: "What is a comet made of?", Answer: "Ice, dust, and rocky material"},
        {Question: "What is the Kuiper Belt?", Answer: "A region beyond Neptune containing icy objects"},
        {Question: "What is Pluto classified as?", Answer: "A dwarf planet"},
    }
    
    for _, fact := range planetFacts {
        if len(additionalFacts) >= needed { break }
        additionalFacts = append(additionalFacts, fact)
    }
    
    moreDiverseFacts := []models.Fact{
        {Question: "What is the most abundant gas in Earth's atmosphere?", Answer: "Nitrogen"},
        {Question: "What percentage of Earth's surface is covered by water?", Answer: "71%"},
        {Question: "What is the pH of pure water?", Answer: "7"},
        {Question: "What is the hardest rock?", Answer: "Diamond"},
        {Question: "What is the most common blood type?", Answer: "O positive"},
        {Question: "What is the rarest blood type?", Answer: "AB negative"},
        {Question: "What is the largest country by land area?", Answer: "Russia"},
        {Question: "What is the smallest continent?", Answer: "Australia"},
        {Question: "What is the longest word in English?", Answer: "Pneumonoultramicroscopicsilicovolcanoconios"},
        {Question: "What is the most spoken language in the world?", Answer: "Mandarin Chinese"},
    }
    
    for len(additionalFacts) < needed && len(moreDiverseFacts) > 0 {
        additionalFacts = append(additionalFacts, moreDiverseFacts...)
        if len(additionalFacts) >= needed {
            additionalFacts = additionalFacts[:needed]
            break
        }
        

        startNum := len(additionalFacts)
        for i := startNum; i < startNum+100 && len(additionalFacts) < needed; i++ {
            additionalFacts = append(additionalFacts, models.Fact{
                Question: fmt.Sprintf("What is %d plus %d?", i%50+1, (i+1)%50+1),
                Answer:   fmt.Sprintf("%d", (i%50+1)+((i+1)%50+1)),
            })
        }
    }
    
    return additionalFacts
}

var facts = generateFacts()

func main() {
    database.ConnectDb()

    facts := generateFacts()
    
    fmt.Printf("Generated %d facts\n", len(facts))

    database.DB.Db.Where("1 = 1").Delete(&models.Fact{})

    batchSize := 100
    for i := 0; i < len(facts); i += batchSize {
        end := i + batchSize
        if end > len(facts) {
            end = len(facts)
        }
        
        batch := facts[i:end]
        result := database.DB.Db.Create(&batch)
        if result.Error != nil {
            log.Printf("Error inserting batch %d-%d: %v", i+1, end, result.Error)
        } else {
            fmt.Printf("Inserted facts %d-%d\n", i+1, end)
        }
    }

    fmt.Printf("Successfully inserted %d facts into the database!\n", len(facts))
}
