### Cointoss, in Go!

A little fun project I made which (pseudo) randomly flips a coin one million times,
displays how many times you got heads or tails and then tells you the difference between the two.

This uses Go's random function from the math library and Time.now on Unix which gets the current time
to generate a random seed for the random function, it basically guarantees it being
random if it's ran a millisecond apart.


### Running on your system
Prerequisites:
- [x] Go
- [x] Git
- [x] An internet connection for cloning the repo.
```bash
git clone https://github.com/AkisaRazbu/Cointoss
cd Cointoss
go run .
```
