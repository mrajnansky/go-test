package main

import (
    "fmt"
    "log"

    "github.com/mrajnansky/go-test/database"
    "github.com/mrajnansky/go-test/models"
)

var facts = []models.Fact{
    {Question: "What is the capital of France?", Answer: "Paris"},
    {Question: "What is the largest planet in our solar system?", Answer: "Jupiter"},
    {Question: "Who painted the Mona Lisa?", Answer: "Leonardo da Vinci"},
    {Question: "What is the chemical symbol for gold?", Answer: "Au"},
    {Question: "What year did World War II end?", Answer: "1945"},
    {Question: "What is the smallest country in the world?", Answer: "Vatican City"},
    {Question: "Who wrote Romeo and Juliet?", Answer: "William Shakespeare"},
    {Question: "What is the hardest natural substance on Earth?", Answer: "Diamond"},
    {Question: "How many continents are there?", Answer: "Seven"},
    {Question: "What is the longest river in the world?", Answer: "The Nile River"},
    {Question: "What gas do plants absorb from the atmosphere?", Answer: "Carbon dioxide"},
    {Question: "What is the speed of light?", Answer: "299,792,458 meters per second"},
    {Question: "Who invented the telephone?", Answer: "Alexander Graham Bell"},
    {Question: "What is the largest ocean on Earth?", Answer: "Pacific Ocean"},
    {Question: "How many bones are in the human body?", Answer: "206"},
    {Question: "What is the currency of Japan?", Answer: "Yen"},
    {Question: "Who composed The Four Seasons?", Answer: "Antonio Vivaldi"},
    {Question: "What is the boiling point of water?", Answer: "100 degrees Celsius"},
    {Question: "What is the largest mammal in the world?", Answer: "Blue whale"},
    {Question: "Who discovered penicillin?", Answer: "Alexander Fleming"},
    {Question: "What is the tallest mountain in the world?", Answer: "Mount Everest"},
    {Question: "How many chambers does a human heart have?", Answer: "Four"},
    {Question: "What is the capital of Australia?", Answer: "Canberra"},
    {Question: "Who wrote Pride and Prejudice?", Answer: "Jane Austen"},
    {Question: "What is the chemical formula for water?", Answer: "H2O"},
    {Question: "What year was the first moon landing?", Answer: "1969"},
    {Question: "What is the smallest unit of matter?", Answer: "Atom"},
    {Question: "Who painted Starry Night?", Answer: "Vincent van Gogh"},
    {Question: "What is the largest desert in the world?", Answer: "Antarctica"},
    {Question: "How many sides does a hexagon have?", Answer: "Six"},
    {Question: "What is the capital of Brazil?", Answer: "Brasília"},
    {Question: "Who invented the light bulb?", Answer: "Thomas Edison"},
    {Question: "What is the fastest land animal?", Answer: "Cheetah"},
    {Question: "How many minutes are in a full day?", Answer: "1,440"},
    {Question: "What is the chemical symbol for iron?", Answer: "Fe"},
    {Question: "Who wrote 1984?", Answer: "George Orwell"},
    {Question: "What is the largest bone in the human body?", Answer: "Femur"},
    {Question: "What planet is known as the Red Planet?", Answer: "Mars"},
    {Question: "How many strings does a violin have?", Answer: "Four"},
    {Question: "What is the capital of Canada?", Answer: "Ottawa"},
    {Question: "Who discovered gravity?", Answer: "Isaac Newton"},
    {Question: "What is the deepest ocean trench?", Answer: "Mariana Trench"},
    {Question: "How many teeth does an adult human have?", Answer: "32"},
    {Question: "What is the currency of the United Kingdom?", Answer: "Pound Sterling"},
    {Question: "Who composed Beethoven's 9th Symphony?", Answer: "Ludwig van Beethoven"},
    {Question: "What is the freezing point of water?", Answer: "0 degrees Celsius"},
    {Question: "What is the largest bird in the world?", Answer: "Ostrich"},
    {Question: "Who invented the World Wide Web?", Answer: "Tim Berners-Lee"},
    {Question: "What is the second largest planet in our solar system?", Answer: "Saturn"},
    {Question: "How many eyes does a spider have?", Answer: "Eight"},
    {Question: "What is the capital of Germany?", Answer: "Berlin"},
    {Question: "Who wrote To Kill a Mockingbird?", Answer: "Harper Lee"},
    {Question: "What is the chemical symbol for oxygen?", Answer: "O"},
    {Question: "What year did the Berlin Wall fall?", Answer: "1989"},
    {Question: "What is the hardest rock?", Answer: "Diamond"},
    {Question: "Who painted The Last Supper?", Answer: "Leonardo da Vinci"},
    {Question: "What is the smallest planet in our solar system?", Answer: "Mercury"},
    {Question: "How many legs does a spider have?", Answer: "Eight"},
    {Question: "What is the capital of Italy?", Answer: "Rome"},
    {Question: "Who invented the airplane?", Answer: "Wright Brothers"},
    {Question: "What is the fastest marine animal?", Answer: "Sailfish"},
    {Question: "How many seconds are in an hour?", Answer: "3,600"},
    {Question: "What is the chemical symbol for silver?", Answer: "Ag"},
    {Question: "Who wrote The Great Gatsby?", Answer: "F. Scott Fitzgerald"},
    {Question: "What is the longest bone in the human body?", Answer: "Femur"},
    {Question: "What planet is closest to the Sun?", Answer: "Mercury"},
    {Question: "How many strings does a guitar have?", Answer: "Six"},
    {Question: "What is the capital of Spain?", Answer: "Madrid"},
    {Question: "Who discovered America?", Answer: "Christopher Columbus"},
    {Question: "What is the highest waterfall in the world?", Answer: "Angel Falls"},
    {Question: "How many ribs does a human have?", Answer: "24"},
    {Question: "What is the currency of China?", Answer: "Yuan"},
    {Question: "Who composed The Magic Flute?", Answer: "Wolfgang Amadeus Mozart"},
    {Question: "What is the melting point of ice?", Answer: "0 degrees Celsius"},
    {Question: "What is the largest fish in the world?", Answer: "Whale shark"},
    {Question: "Who invented the printing press?", Answer: "Johannes Gutenberg"},
    {Question: "What is the third planet from the Sun?", Answer: "Earth"},
    {Question: "How many wings does a bee have?", Answer: "Four"},
    {Question: "What is the capital of Russia?", Answer: "Moscow"},
    {Question: "Who wrote Harry Potter?", Answer: "J.K. Rowling"},
    {Question: "What is the chemical symbol for carbon?", Answer: "C"},
    {Question: "What year did the Titanic sink?", Answer: "1912"},
    {Question: "What is the softest mineral?", Answer: "Talc"},
    {Question: "Who painted The Scream?", Answer: "Edvard Munch"},
    {Question: "What is the hottest planet in our solar system?", Answer: "Venus"},
    {Question: "How many tentacles does an octopus have?", Answer: "Eight"},
    {Question: "What is the capital of Egypt?", Answer: "Cairo"},
    {Question: "Who invented the steam engine?", Answer: "James Watt"},
    {Question: "What is the slowest mammal?", Answer: "Sloth"},
    {Question: "How many days are in a leap year?", Answer: "366"},
    {Question: "What is the chemical symbol for hydrogen?", Answer: "H"},
    {Question: "Who wrote Lord of the Rings?", Answer: "J.R.R. Tolkien"},
    {Question: "What is the smallest muscle in the human body?", Answer: "Stapedius"},
    {Question: "What planet has the most moons?", Answer: "Jupiter"},
    {Question: "How many chambers does a fish heart have?", Answer: "Two"},
    {Question: "What is the capital of India?", Answer: "New Delhi"},
    {Question: "Who discovered electricity?", Answer: "Benjamin Franklin"},
    {Question: "What is the largest lake in the world?", Answer: "Caspian Sea"},
    {Question: "How many vertebrae are in the human spine?", Answer: "33"},
    {Question: "What is the currency of Mexico?", Answer: "Peso"},
    {Question: "Who composed Carmen?", Answer: "Georges Bizet"},
    {Question: "What is the sublimation point of dry ice?", Answer: "-78.5 degrees Celsius"},
    {Question: "What is the smallest dog breed?", Answer: "Chihuahua"},
    {Question: "Who invented the radio?", Answer: "Guglielmo Marconi"},
}

func main() {
    // Connect to database
    database.ConnectDb()

    // Clear existing facts (optional)
    database.DB.Db.Where("1 = 1").Delete(&models.Fact{})

    // Insert facts
    for i, fact := range facts {
        result := database.DB.Db.Create(&fact)
        if result.Error != nil {
            log.Printf("Error inserting fact %d: %v", i+1, result.Error)
        } else {
            fmt.Printf("Inserted fact %d: %s\n", i+1, fact.Question)
        }
    }

    fmt.Printf("Successfully inserted %d facts into the database!\n", len(facts))
}
