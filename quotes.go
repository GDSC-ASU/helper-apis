package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func handleGetQuote(w http.ResponseWriter, r *http.Request) {
	rand := rand.New(rand.NewSource(time.Now().UnixMilli()))
	quote := quotes[rand.Intn(len(quotes))]
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(quote)
}

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []Quote{
	{
		Text:   "The secret of joy in work is contained in one word � excellence. To know how to do something well is to enjoy it.",
		Author: "Pearl Buck",
	},
	{
		Text:   "To be happy is to be able to become aware of oneself without fright.",
		Author: "Walter Benjamin",
	},
	{
		Text:   "Into each life rain must fall but rain can be the giver of life and it is all in your attitude that makes rain produce sunshine.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "True silence is the rest of the mind; it is to the spirit what sleep is to the body, nourishment and refreshment.",
		Author: "William Penn",
	},
	{
		Text:   "To give hope to someone occurs when you teach them how to use the tools to do it for themselves.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "Ability will never catch up with the demand for it.",
		Author: "Confucius",
	},
	{
		Text:   "If one advances confidently in the direction of his dream, and endeavours to live the life which he had imagines, he will meet with a success unexpected in common hours.",
		Author: "Henry David Thoreau",
	},
	{
		Text:   "Action will remove the doubts that theory cannot solve.",
		Author: "Tehyi Hsieh",
	},
	{
		Text:   "Kindness is the golden chain by which society is bound together.",
		Author: "Johann Wolfgang von Goethe",
	},
	{
		Text:   "To accomplish great things, we must dream as well as act.",
		Author: "Anatole France",
	},
	{
		Text:   "Never bend your head. Always hold it high. Look the world right in the eye.",
		Author: "Helen Keller",
	},
	{Text: "Take it easy � but take it.", Author: "Woody Guthrie"},
	{
		Text:   "They say that time changes things, but you actually have to change them yourself.",
		Author: "Andy Warhol",
	},
	{Text: "He who knows himself is enlightened.", Author: "Lao Tzu"},
	{
		Text:   "Every sixty seconds you spend angry, upset or mad, is a full minute of happiness you�ll never get back.",
		Author: "Someone",
	},
	{
		Text:   "Being in humaneness is good. If we select other goodness and thus are far apart from humaneness, how can we be the wise?",
		Author: "Confucius",
	},
	{
		Text:   "A true friend is the most precious of all possessions and the one we take the least thought about acquiring.",
		Author: "Francois de La Rochefoucauld",
	},
	{
		Text:   "Giving up doesn't always mean you are weak. Sometimes it means that you are strong enough to let go.",
		Author: "Someone",
	},
	{
		Text:   "God has given you one face, and you make yourself another.",
		Author: "William Shakespeare",
	},
	{
		Text:   "Though no one can go back and make a brand new start, anyone can start from now and make a brand new ending.",
		Author: "Someone",
	},
	{
		Text:   "The fox has many tricks. The hedgehog has but one. But that is the best of all.",
		Author: "Desiderius Erasmus",
	},
	{
		Text:   "Every adversity, every failure, every heartache carries with it the seed of an equal or greater benefit.",
		Author: "Napoleon Hill",
	},
	{
		Text:   "If a man does his best, what else is there?",
		Author: "George Patton",
	},
	{
		Text:   "Your outlook on life is a direct reflection on how much you like yourself.",
		Author: "Lululemon",
	},
	{
		Text:   "The greatest remedy for anger is delay.",
		Author: "Seneca",
	},
	{
		Text:   "What lies behind us and what lies before us are tiny matters compared to what lies within us.",
		Author: "Walt Emerson",
	},
	{
		Text:   "Ability will never catch up with the demand for it.",
		Author: "Confucius",
	},
	{
		Text:   "Doing nothing is better than being busy doing nothing.",
		Author: "Lao Tzu",
	},
	{
		Text:   "Happiness is not something ready made. It comes from your own actions.",
		Author: "Dalai Lama",
	},
	{
		Text:   "Everyone is a genius at least once a year. A real genius has his original ideas closer together.",
		Author: "Georg Lichtenberg",
	},
	{
		Text:   "Moral excellence comes about as a result of habit. We become just by doing just acts, temperate by doing temperate acts, brave by doing brave acts.",
		Author: "Aristotle",
	},
	{Text: "A jug fills drop by drop.", Author: "Buddha"},
	{
		Text:   "The best way to pay for a lovely moment is to enjoy it.",
		Author: "Richard Bach",
	},
	{
		Text:   "Men of perverse opinion do not know the excellence of what is in their hands, till some one dash it from them.",
		Author: "Sophocles",
	},
	{
		Text:   "In the long run we get no more than we have been willing to risk giving.",
		Author: "Sheldon Kopp",
	},
	{
		Text:   "Liberty, taking the word in its concrete sense, consists in the ability to choose.",
		Author: "Simone Weil",
	},
	{
		Text:   "You have enemies? Good. That means you've stood up for something, sometime in your life.",
		Author: "Winston Churchill",
	},
	{
		Text:   "Be slow of tongue and quick of eye.",
		Author: "Cervantes",
	},
	{
		Text:   "Great ideas often receive violent opposition from mediocre minds.",
		Author: "Albert Einstein",
	},
	{
		Text:   "Don't smother each other. No one can grow in the shade.",
		Author: "Leo F. Buscaglia",
	},
	{
		Text:   "A failure is a man who has blundered but is not capable of cashing in on the experience.",
		Author: "Elbert Hubbard",
	},
	{
		Text:   "Action is the foundational key to all success.",
		Author: "Pablo Picasso",
	},
	{
		Text:   "Argue for your limitations, and sure enough theyre yours.",
		Author: "Richard Bach",
	},
	{
		Text:   "Nothing happens unless first we dream.",
		Author: "Carl Sandburg",
	},
	{
		Text:   "An optimist is a person who sees a green light everywhere, while the pessimist sees only the red spotlight... The truly wise person is colour-blind.",
		Author: "Albert Schweitzer",
	},
	{
		Text:   "Nothing is softer or more flexible than water, yet nothing can resist it.",
		Author: "Lao Tzu",
	},
	{
		Text:   "Don't smother each other. No one can grow in the shade.",
		Author: "Leo F. Buscaglia",
	},
	{
		Text:   "We are all faced with a series of great opportunities brilliantly disguised as impossible situations.",
		Author: "Charles R. Swindoll",
	},
	{
		Text:   "The real measure of your wealth is how much youd be worth if you lost all your money.",
		Author: "Someone",
	},
	{
		Text:   "Never say there is nothing beautiful in the world any more. There is always something to make you wonder in the shape of a tree, the trembling of a leaf.",
		Author: "Albert Schweitzer",
	},
	{
		Text:   "There is no need for temples, no need for complicated philosophies. My brain and my heart are my temples; my philosophy is kindness.",
		Author: "Dalai Lama",
	},
	{
		Text:   "If you seek truth you will not seek victory by dishonourable means, and if you find truth you will become invincible.",
		Author: "Epictetus",
	},
	{
		Text:   "Don't judge each day by the harvest you reap but by the seeds you plant.",
		Author: "Robert Stevenson",
	},
	{
		Text:   "Nothing happens unless first we dream.",
		Author: "Carl Sandburg",
	},
	{
		Text:   "The truth is always exciting. Speak it, then. Life is dull without it.",
		Author: "Pearl Buck",
	},
	{
		Text:   "If you love someone, set them free. If they come back they're yours; if they don't they never were.",
		Author: "Richard Bach",
	},
	{
		Text:   "By going beyond your own problems and taking care of others, you gain inner strength, self-confidence, courage, and a greater sense of calm.",
		Author: "Dalai Lama",
	},
	{
		Text:   "Watch the little things; a small leak will sink a great ship.",
		Author: "Benjamin Franklin",
	},
	{
		Text:   "Never idealize others. They will never live up to your expectations.",
		Author: "Leo Buscaglia",
	},
	{
		Text:   "You have to take it as it happens, but you should try to make it happen the way you want to take it.",
		Author: "Old German proverb",
	},
	{
		Text:   "Notice that the stiffest tree is most easily cracked, while the bamboo or willow survives by bending with the wind.",
		Author: "Bruce Lee",
	},
	{
		Text:   "Action may not always bring happiness, but there is no happiness without action.",
		Author: "Benjamin Disraeli",
	},
	{
		Text:   "Responsibility is not inherited, it is a choice that everyone needs to make at some point in their life.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "You will not be punished for your anger, you will be punished by your anger.",
		Author: "Buddha",
	},
	{
		Text:   "The greatest part of our happiness depends on our dispositions, not our circumstances.",
		Author: "Martha Washington",
	},
	{
		Text:   "The smallest flower is a thought, a life answering to some feature of the Great Whole, of whom they have a persistent intuition.",
		Author: "Honore de Balzac",
	},
	{
		Text:   "Face your deficiencies and acknowledge them; but do not let them master you. Let them teach you patience, sweetness, insight.",
		Author: "Helen Keller",
	},
	{
		Text:   "Before you put on a frown, make absolutely sure there are no smiles available.",
		Author: "Jim Beggs",
	},
	{
		Text:   "Those who dream by day are cognizant of many things which escape those who dream only by night.",
		Author: "Edgar Allan Poe",
	},
	{
		Text:   "Being in humaneness is good. If we select other goodness and thus are far apart from humaneness, how can we be the wise?",
		Author: "Confucius",
	},
	{
		Text:   "Life is what you make of it. Always has been, always will be.",
		Author: "Grandma Moses",
	},
	{
		Text:   "One today is worth two tomorrows.",
		Author: "Benjamin Franklin",
	},
	{
		Text:   "Every problem has a gift for you in its hands.",
		Author: "Richard Bach",
	},
	{
		Text:   "More often than not, anger is actually an indication of weakness rather than of strength.",
		Author: "Dalai Lama",
	},
	{
		Text:   "The day you decide to do it is your lucky day.",
		Author: "Japanese proverb",
	},
	{
		Text:   "Success is not the key to happiness. Happiness is the key to success. If you love what you are doing, you will be successful.",
		Author: "Albert Schweitzer",
	},
	{
		Text:   "A fine quotation is a diamond on the finger of a man of wit, and a pebble in the hand of a fool.",
		Author: "Joseph Roux",
	},
	{
		Text:   "Inspiration exists, but it has to find us working.",
		Author: "Pablo Picasso",
	},
	{
		Text:   "I have just three things to teach: simplicity, patience, compassion. These three are your greatest treasures.",
		Author: "Lao Tzu",
	},
	{
		Text:   "There are two primary choices in life: to accept conditions as they exist, or accept the responsibility for changing them.",
		Author: "Denis Waitley",
	},
	{
		Text:   "Neither genius, fame, nor love show the greatness of the soul. Only kindness can do that.",
		Author: "Jean Lacordaire",
	},
	{Text: "What we think, we become.", Author: "Buddha"},
	{
		Text:   "Work for something because it is good, not just because it stands a chance to succeed.",
		Author: "Vaclav Havel",
	},
	{
		Text:   "By letting it go it all gets done. The world is won by those who let it go. But when you try and try. The world is beyond the winning.",
		Author: "Lao Tzu",
	},
	{
		Text:   "The more you know yourself, the more you forgive yourself.",
		Author: "Confucius",
	},
	{
		Text:   "It is better to understand a little than to misunderstand a lot.",
		Author: "Anatole France",
	},
	{
		Text:   "If one is estranged from oneself, then one is estranged from others too. If one is out of touch with oneself, then one cannot touch others.",
		Author: "Anne Lindbergh",
	},
	{
		Text:   "Each day provides its own gifts.",
		Author: "Marcus Aurelius",
	},
	{
		Text:   "No man can succeed in a line of endeavor which he does not like.",
		Author: "Napoleon Hill",
	},
	{
		Text:   "Action is the foundational key to all success.",
		Author: "Pablo Picasso",
	},
	{
		Text:   "Through meditation and by giving full attention to one thing at a time, we can learn to direct attention where we choose.",
		Author: "Eknath Easwaran",
	},
	{
		Text:   "People grow through experience if they meet life honestly and courageously. This is how character is built.",
		Author: "Eleanor Roosevelt",
	},
	{
		Text:   "Well done is better than well said.",
		Author: "Benjamin Franklin",
	},
	{
		Text:   "I never think of the future. It comes soon enough.",
		Author: "Albert Einstein",
	},
	{
		Text:   "Practice yourself, for heavens sake in little things, and then proceed to greater.",
		Author: "Epictetus",
	},
	{
		Text:   "You don't drown by falling in water. You drown by staying there.",
		Author: "Someone",
	},
	{
		Text:   "First comes thought; then organization of that thought, into ideas and plans; then transformation of those plans into reality. The beginning, as you will observe, is in your imagination.",
		Author: "Napoleon Hill",
	},
	{
		Text:   "All great achievements require time.",
		Author: "Maya Angelou",
	},
	{
		Text:   "Difficulties are meant to rouse, not discourage. The human spirit is to grow strong by conflict.",
		Author: "William Channing",
	},
	{
		Text:   "A failure is a man who has blundered but is not capable of cashing in on the experience.",
		Author: "Elbert Hubbard",
	},
	{
		Text:   "The world is but a canvas to the imagination.",
		Author: "Henry Thoreau",
	},
	{
		Text:   "The best cure for the body is a quiet mind.",
		Author: "Napoleon Bonaparte",
	},
	{
		Text:   "The final proof of greatness lies in being able to endure criticism without resentment.",
		Author: "Elbert Hubbard",
	},
	{
		Text:   "What is not started today is never finished tomorrow.",
		Author: "Goethe",
	},
	{
		Text:   "You, yourself, as much as anybody in the entire universe, deserve your love and affection.",
		Author: "Buddha",
	},
	{
		Text:   "Until you value yourself, you won't value your time. Until you value your time, you won't do anything with it.",
		Author: "M. Scott Peck",
	},
	{
		Text:   "Freedom is what you do with what's been done to you.",
		Author: "Jean-Paul Sartre",
	},
	{
		Text:   "What you fear is that which requires action to overcome.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "God has given you one face, and you make yourself another.",
		Author: "William Shakespeare",
	},
	{
		Text:   "Everyone smiles in the same language.",
		Author: "Someone",
	},
	{
		Text:   "He that never changes his opinions, never corrects his mistakes, and will never be wiser on the morrow than he is today.",
		Author: "Tryon Edwards",
	},
	{
		Text:   "The awareness of our own strength makes us modest.",
		Author: "Paul Cezanne",
	},
	{
		Text:   "Truth is powerful and it prevails.",
		Author: "Sojourner Truth",
	},
	{
		Text:   "Of course there is no formula for success except perhaps an unconditional acceptance of life and what it brings.",
		Author: "Arthur Rubinstein",
	},
	{
		Text:   "I always wanted to be somebody, but I should have been more specific.",
		Author: "Lily Tomlin",
	},
	{
		Text:   "Cherish your visions and your dreams as they are the children of your soul, the blueprints of your ultimate achievements.",
		Author: "Napoleon Hill",
	},
	{
		Text:   "Being right is highly overrated. Even a stopped clock is right twice a day.",
		Author: "Someone",
	},
	{
		Text:   "Friends are those rare people who ask how we are and then wait to hear the answer.",
		Author: "Ed Cunningham",
	},
	{
		Text:   "The minute you settle for less than you deserve, you get even less than you settled for.",
		Author: "Maureen Dowd",
	},
	{
		Text:   "What you do not want done to yourself, do not do to others.",
		Author: "Confucius",
	},
	{
		Text:   "The foolish man seeks happiness in the distance, the wise grows it under his feet.",
		Author: "James Oppenheim",
	},
	{
		Text:   "However many holy words you read, However many you speak, What good will they do you If you do not act on upon them?",
		Author: "Buddha",
	},
	{
		Text:   "When you see a man of worth, think of how you may emulate him. When you see one who is unworthy, examine yourself.",
		Author: "Confucius",
	},
	{Text: "Talk doesn't cook rice.", Author: "Chinese proverb"},
	{
		Text:   "Four steps to achievement: Plan purposefully. Prepare prayerfully. Proceed positively. Pursue persistently.",
		Author: "William Arthur Ward",
	},
	{
		Text:   "From wonder into wonder existence opens.",
		Author: "Lao Tzu",
	},
	{
		Text:   "As long as your going to be thinking anyway, think big.",
		Author: "Donald Trump",
	},
	{
		Text:   "The road leading to a goal does not separate you from the destination; it is essentially a part of it.",
		Author: "Charles DeLint",
	},
	{
		Text:   "Work out your own salvation. Do not depend on others.",
		Author: "Buddha",
	},
	{
		Text:   "The best and most beautiful things in the world cannot be seen, nor touched... but are felt in the heart.",
		Author: "Helen Keller",
	},
	{
		Text:   "You can't let praise or criticism get to you. It's a weakness to get caught up in either one.",
		Author: "John Wooden",
	},
	{
		Text:   "The journey of a thousand miles begins with one step.",
		Author: "Lao Tzu",
	},
	{Text: "What we think, we become.", Author: "Buddha"},
	{
		Text:   "I'm not afraid of storms, for Im learning how to sail my ship.",
		Author: "Louisa Alcott",
	},
	{
		Text:   "Life is not measured by the breaths you take, but by its breathtaking moments.",
		Author: "Michael Vance",
	},
	{
		Text:   "Fate is in your hands and no one elses",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "From small beginnings come great things.",
		Author: "Someone",
	},
	{
		Text:   "Victory belongs to the most persevering.",
		Author: "Napoleon Bonaparte",
	},
	{
		Text:   "As our case is new, we must think and act anew.",
		Author: "Abraham Lincoln",
	},
	{
		Text:   "If we have a positive mental attitude, then even when surrounded by hostility, we shall not lack inner peace.",
		Author: "Dalai Lama",
	},
	{
		Text:   "Know, first, who you are, and then adorn yourself accordingly.",
		Author: "Epictetus",
	},
	{
		Text:   "Those who are free of resentful thoughts surely find peace.",
		Author: "Buddha",
	},
	{
		Text:   "Every great advance in science has issued from a new audacity of the imagination.",
		Author: "John Dewey",
	},
	{
		Text:   "Every gift from a friend is a wish for your happiness.",
		Author: "Richard Bach",
	},
	{
		Text:   "If you cannot be silent be brilliant and thoughtful.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "Beware of the half truth. You may have gotten hold of the wrong half.",
		Author: "Someone",
	},
	{
		Text:   "Smile, breathe, and go slowly.",
		Author: "Thich Nhat Hanh",
	},
	{
		Text:   "Always be yourself, express yourself, have faith in yourself, do not go out and look for a successful personality and duplicate it.",
		Author: "Bruce Lee",
	},
	{
		Text:   "Meaning is not what you start with but what you end up with.",
		Author: "Peter Elbow",
	},
	{
		Text:   "Make the most of yourself for that is all there is of you.",
		Author: "Ralph Emerson",
	},
	{
		Text:   "The only limit to your impact is your imagination and commitment.",
		Author: "Tony Robbins",
	},
	{
		Text:   "The greatest good you can do for another is not just to share your riches but to reveal to him his own.",
		Author: "Benjamin Disraeli",
	},
	{
		Text:   "Each day provides its own gifts.",
		Author: "Marcus Aurelius",
	},
	{
		Text:   "A little more persistence, a little more effort, and what seemed hopeless failure may turn to glorious success.",
		Author: "Elbert Hubbard",
	},
	{
		Text:   "If we have a positive mental attitude, then even when surrounded by hostility, we shall not lack inner peace.",
		Author: "Dalai Lama",
	},
	{
		Text:   "Ability will never catch up with the demand for it.",
		Author: "Confucius",
	},
	{
		Text:   "Forgiveness does not change the past, but it does enlarge the future.",
		Author: "Paul Boese",
	},
	{
		Text:   "The greatest danger for most of us is not that our aim is too high and we miss it, but that it is too low and we reach it.",
		Author: "Michelangelo",
	},
	{
		Text:   "Sunshine is delicious, rain is refreshing, wind braces us up, snow is exhilarating; there is really no such thing as bad weather, only different kinds of good weather.",
		Author: "John Ruskin",
	},
	{
		Text:   "Better than a thousand hollow words, is one word that brings peace.",
		Author: "Buddha",
	},
	{
		Text:   "When you don't know what you believe, everything becomes an argument. Everything is debatable. But when you stand for something, decisions are obvious.",
		Author: "Someone",
	},
	{
		Text:   "From error to error one discovers the entire truth.",
		Author: "Sigmund Freud",
	},
	{
		Text:   "If there is no struggle, there is no progress.",
		Author: "Frederick Douglass",
	},
	{
		Text:   "It all depends on how we look at things, and not how they are in themselves.",
		Author: "Carl Jung",
	},
	{
		Text:   "Take heed: you do not find what you do not seek.",
		Author: "English proverb",
	},
	{
		Text:   "Give it all you've got because you never know if there's going to be a next time.",
		Author: "Danielle Ingrum",
	},
	{
		Text:   "No day in which you learn something is a complete loss.",
		Author: "David Eddings",
	},
	{
		Text:   "Reviewing what you have learned and learning anew, you are fit to be a teacher.",
		Author: "Confucius",
	},
	{
		Text:   "Sadness flies away on the wings of time.",
		Author: "Jean de la Fontaine",
	},
	{
		Text:   "I love my past. I love my present. Im not ashamed of what Ive had, and Im not sad because I have it no longer.",
		Author: "Colette",
	},
	{
		Text:   "One of the advantages of being disorderly is that one is constantly making exciting discoveries.",
		Author: "A. A. Milne",
	},
	{
		Text:   "The final proof of greatness lies in being able to endure criticism without resentment.",
		Author: "Elbert Hubbard",
	},
	{
		Text:   "Promises are the uniquely human way of ordering the future, making it predictable and reliable to the extent that this is humanly possible.",
		Author: "Hannah Arendt",
	},
	{
		Text:   "Better to have loved and lost, than to have never loved at all.",
		Author: "St. Augustine",
	},
	{
		Text:   "An invasion of armies can be resisted, but not an idea whose time has come.",
		Author: "Victor Hugo",
	},
	{
		Text:   "Whatever happens, take responsibility.",
		Author: "Tony Robbins",
	},
	{
		Text:   "The moment one gives close attention to anything, it becomes a mysterious, awesome, indescribably magnificent world in itself.",
		Author: "Henry Miller",
	},
	{
		Text:   "Cherish your visions and your dreams as they are the children of your soul, the blueprints of your ultimate achievements.",
		Author: "Napoleon Hill",
	},
	{
		Text:   "This world, after all our science and sciences, is still a miracle; wonderful, inscrutable, magical and more, to whosoever will think of it.",
		Author: "Thomas Carlyle",
	},
	{
		Text:   "You block your dream when you allow your fear to grow bigger than your faith.",
		Author: "Mary Morrissey",
	},
	{
		Text:   "We can only learn to love by loving.",
		Author: "Iris Murdoch",
	},
	{
		Text:   "The best thing in every noble dream is the dreamer...",
		Author: "Moncure Conway",
	},
	{
		Text:   "Opportunity often comes disguised in the form of misfortune, or temporary defeat.",
		Author: "Napoleon Hill",
	},
	{
		Text:   "Don't be dismayed by good-byes. A farewell is necessary before you can meet again. And meeting again, after moments or lifetimes, is certain for those who are friends.",
		Author: "Richard Bach",
	},
	{
		Text:   "The truth you believe and cling to makes you unavailable to hear anything new.",
		Author: "Pema Chodron",
	},
	{
		Text:   "Though no one can go back and make a brand new start, anyone can start from now and make a brand new ending.",
		Author: "Someone",
	},
	{
		Text:   "Meaning is not what you start with but what you end up with.",
		Author: "Peter Elbow",
	},
	{
		Text:   "To be able to give away riches is mandatory if you wish to possess them. This is the only way that you will be truly rich.",
		Author: "Mahummad Ali",
	},
	{
		Text:   "Life shrinks or expands in proportion to one's courage.",
		Author: "Anais Nin",
	},
	{
		Text:   "In seed time learn, in harvest teach, in winter enjoy.",
		Author: "William Blake",
	},
	{
		Text:   "If you spend your whole life waiting for the storm, you'll never enjoy the sunshine.",
		Author: "Morris West",
	},
	{
		Text:   "Thousands of candles can be lit from a single, and the life of the candle will not be shortened. Happiness never decreases by being shared.",
		Author: "Buddha",
	},
	{
		Text:   "Don't ruin the present with the ruined past.",
		Author: "Ellen Gilchrist",
	},
	{
		Text:   "Love is the master key that opens the gates of happiness.",
		Author: "Oliver Holmes",
	},
	{
		Text:   "The superior man acts before he speaks, and afterwards speaks according to his action.",
		Author: "Confucius",
	},
	{
		Text:   "If A is success in life, then A equals x plus y plus z. Work is x; y is play; and z is keeping your mouth shut.",
		Author: "Albert Einstein",
	},
	{
		Text:   "The least of things with a meaning is worth more in life than the greatest of things without it.",
		Author: "Carl Jung",
	},
	{
		Text:   "Better be ignorant of a matter than half know it.",
		Author: "Publilius Syrus",
	},
	{
		Text:   "Arriving at one point is the starting point to another.",
		Author: "John Dewey",
	},
	{Text: "The world is always in movement.", Author: "V. Naipaul"},
	{
		Text:   "When performance exceeds ambition, the overlap is called success.",
		Author: "Cullen Hightower",
	},
	{
		Text:   "The greatest pleasure I know is to do a good action by stealth, and to have it found out by accident.",
		Author: "Charles Lamb",
	},
	{
		Text:   "If we could learn to like ourselves, even a little, maybe our cruelties and angers might melt away.",
		Author: "John Steinbeck",
	},
	{
		Text:   "Life a culmination of the past, an awareness of the present, an indication of the future beyond knowledge, the quality that gives a touch of divinity to matter.",
		Author: "Charles A. Lindbergh",
	},
	{
		Text:   "The only real valuable thing is intuition.",
		Author: "Albert Einstein",
	},
	{
		Text:   "Kind words do not cost much. Yet they accomplish much.",
		Author: "Blaise Pascal",
	},
	{
		Text:   "Don't be pushed by your problems; be led by your dreams.",
		Author: "Someone",
	},
	{
		Text:   "Love is the flower you've got to let grow.",
		Author: "John Lennon",
	},
	{
		Text:   "Most folks are as happy as they make up their minds to be.",
		Author: "Abraham Lincoln",
	},
	{
		Text:   "He that never changes his opinions, never corrects his mistakes, and will never be wiser on the morrow than he is today.",
		Author: "Tryon Edwards",
	},
	{
		Text:   "Smile, breathe, and go slowly.",
		Author: "Thich Nhat Hanh",
	},
	{
		Text:   "Never bend your head. Always hold it high. Look the world right in the eye.",
		Author: "Helen Keller",
	},
	{
		Text:   "There is nothing in a caterpillar that tells you it's going to be a butterfly.",
		Author: "Buckminster Fuller",
	},
	{
		Text:   "If you think you can, you can. And if you think you can't, you're right.",
		Author: "Henry Ford",
	},
	{
		Text:   "I'm not afraid of storms, for Im learning how to sail my ship.",
		Author: "Louisa Alcott",
	},
	{
		Text:   "In the sky, there is no distinction of east and west; people create distinctions out of their own minds and then believe them to be true.",
		Author: "Buddha",
	},
	{
		Text:   "Until you make peace with who you are, you'll never be content with what you have.",
		Author: "Doris Mortman",
	},
	{
		Text:   "Action may not always bring happiness, but there is no happiness without action.",
		Author: "Benjamin Disraeli",
	},
	{
		Text:   "A good plan today is better than a perfect plan tomorrow.",
		Author: "Someone",
	},
	{
		Text:   "Smile, breathe, and go slowly.",
		Author: "Thich Nhat Hanh",
	},
	{
		Text:   "We can only be said to be alive in those moments when our hearts are conscious of our treasures.",
		Author: "Thornton Wilder",
	},
	{
		Text:   "Do not turn back when you are just at the goal.",
		Author: "Publilius Syrus",
	},
	{
		Text:   "Keep your eyes on the stars and your feet on the ground.",
		Author: "Theodore Roosevelt",
	},
	{
		Text:   "When you are content to be simply yourself and don't compare or compete, everybody will respect you.",
		Author: "Laozi",
	},
	{
		Text:   "Freedom is what you do with what's been done to you.",
		Author: "Jean-Paul Sartre",
	},
	{Text: "All things change; nothing perishes.", Author: "Ovid"},
	{
		Text:   "Courage is what it takes to stand up and speak; courage is also what it takes to sit down and listen.",
		Author: "Winston Churchill",
	},
	{
		Text:   "It is not uncommon for people to spend their whole life waiting to start living.",
		Author: "Eckhart Tolle",
	},
	{
		Text:   "Liberty, taking the word in its concrete sense, consists in the ability to choose.",
		Author: "Simone Weil",
	},
	{
		Text:   "He who fears being conquered is sure of defeat.",
		Author: "Napoleon Bonaparte",
	},
	{
		Text:   "Arrogance and rudeness are training wheels on the bicycle of life � for weak people who cannot keep their balance without them.",
		Author: "Laura Teresa Marquez",
	},
	{
		Text:   "Living at risk is jumping off the cliff and building your wings on the way down.",
		Author: "Ray Bradbury",
	},
	{
		Text:   "Let us resolve to be masters, not the victims, of our history, controlling our own destiny without giving way to blind suspicions and emotions.",
		Author: "John Kennedy",
	},
	{
		Text:   "How far that little candle throws its beams! So shines a good deed in a naughty world.",
		Author: "William Shakespeare",
	},
	{
		Text:   "Well done is better than well said.",
		Author: "Benjamin Franklin",
	},
	{
		Text:   "If you don't design your own life plan, chances are you'll fall into someone else's plan. And guess what they have planned for you? Not much.",
		Author: "Jim Rohn",
	},
	{
		Text:   "You were not born a winner, and you were not born a loser. You are what you make yourself be.",
		Author: "Lou Holtz",
	},
	{
		Text:   "Smile, breathe, and go slowly.",
		Author: "Thich Nhat Hanh",
	},
	{
		Text:   "When it is obvious that the goals cannot be reached, don't adjust the goals, adjust the action steps.",
		Author: "Confucius",
	},
	{
		Text:   "What you give is what you get.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "Forget about all the reasons why something may not work. You only need to find one good reason why it will.",
		Author: "Robert Anthony",
	},
	{
		Text:   "If you are patient in one moment of anger, you will escape one hundred days of sorrow.",
		Author: "Chinese proverb",
	},
	{
		Text:   "There is only one success � to be able to spend your life in your own way.",
		Author: "Christopher Morley",
	},
	{
		Text:   "The greatest way to live with honour in this world is to be what we pretend to be.",
		Author: "Socrates",
	},
	{
		Text:   "Life is the flower for which love is the honey.",
		Author: "Victor Hugo",
	},
	{
		Text:   "We are all something, but none of us are everything.",
		Author: "Blaise Pascal",
	},
	{
		Text:   "The personal life deeply lived always expands into truths beyond itself.",
		Author: "Anais Nin",
	},
	{
		Text:   "I believe that we are solely responsible for our choices, and we have to accept the consequences of every deed, word, and thought throughout our lifetime.",
		Author: "Elisabeth Kubler-Ross",
	},
	{
		Text:   "Nothing in life is to be feared. It is only to be understood.",
		Author: "Marie Curie",
	},
	{
		Text:   "If you love someone, set them free. If they come back they're yours; if they don't they never were.",
		Author: "Richard Bach",
	},
	{
		Text:   "Do not turn back when you are just at the goal.",
		Author: "Publilius Syrus",
	},
	{
		Text:   "Wicked people are always surprised to find ability in those that are good.",
		Author: "Marquis Vauvenargues",
	},
	{
		Text:   "If you have made mistakes, there is always another chance for you. You may have a fresh start any moment you choose.",
		Author: "Mary Pickford",
	},
	{
		Text:   "We aim above the mark to hit the mark.",
		Author: "Ralph Emerson",
	},
	{
		Text:   "Of course there is no formula for success except perhaps an unconditional acceptance of life and what it brings.",
		Author: "Arthur Rubinstein",
	},
	{
		Text:   "To climb steep hills requires a slow pace at first.",
		Author: "William Shakespeare",
	},
	{
		Text:   "Today is the tomorrow we worried about yesterday.",
		Author: "Someone",
	},
	{
		Text:   "What is new in the world? Nothing. What is old in the world? Nothing. Everything has always been and will always be.",
		Author: "Sai Baba",
	},
	{
		Text:   "However many holy words you read, however many you speak, what good will they do you if you do not act on upon them?",
		Author: "Buddha",
	},
	{
		Text:   "Everything is perfect in the universe � even your desire to improve it.",
		Author: "Wayne Dyer",
	},
	{
		Text:   "Complaining doesn't change a thing only taking action does.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "Take rest; a field that has rested gives a bountiful crop.",
		Author: "Ovid",
	},
	{
		Text:   "To be aware of a single shortcoming in oneself is more useful than to be aware of a thousand in someone else.",
		Author: "Tenzin Gyatso",
	},
	{Text: "A good rest is half the work.", Author: "Someone"},
	{
		Text:   "By going beyond your own problems and taking care of others, you gain inner strength, self-confidence, courage, and a greater sense of calm.",
		Author: "Dalai Lama",
	},
	{
		Text:   "It is more important to know where you are going than to get there quickly. Do not mistake activity for achievement.",
		Author: "Mabel Newcomber",
	},
	{
		Text:   "Each day can be one of triumph if you keep up your interests.",
		Author: "George Matthew Adams",
	},
	{
		Text:   "The personal life deeply lived always expands into truths beyond itself.",
		Author: "Anais Nin",
	},
	{
		Text:   "The person who makes a success of living is the one who see his goal steadily and aims for it unswervingly. That is dedication.",
		Author: "Cecil B. DeMille",
	},
	{
		Text:   "Time you enjoy wasting, was not wasted.",
		Author: "John Lennon",
	},
	{
		Text:   "Fear of failure is one attitude that will keep you at the same point in your life.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "Be great in act, as you have been in thought.",
		Author: "William Shakespeare",
	},
	{
		Text:   "If you let go a little, you will have a little peace. If you let go a lot, you will have a lot of peace.",
		Author: "Ajahn Chah",
	},
	{
		Text:   "Setting an example is not the main means of influencing another, it is the only means.",
		Author: "Albert Einstein",
	},
	{
		Text:   "Be miserable. Or motivate yourself. Whatever has to be done, it's always your choice.",
		Author: "Wayne Dyer",
	},
	{
		Text:   "We must not allow ourselves to become like the system we oppose.",
		Author: "Bishop Desmond Tutu",
	},
	{
		Text:   "There are no limitations to the mind except those we acknowledge.",
		Author: "Napoleon Hill",
	},
	{
		Text:   "The more you care, the stronger you can be.",
		Author: "Jim Rohn",
	},
	{
		Text:   "The beginning of wisdom is found in doubting; by doubting we come to the question, and by seeking we may come upon the truth.",
		Author: "Pierre Abelard",
	},
	{
		Text:   "The Creator has not given you a longing to do that which you have no ability to do.",
		Author: "Orison Marden",
	},
	{
		Text:   "Make the best use of what is in your power, and take the rest as it happens.",
		Author: "Epictetus",
	},
	{
		Text:   "Thousands of candles can be lit from a single, and the life of the candle will not be shortened. Happiness never decreases by being shared.",
		Author: "Buddha",
	},
	{
		Text:   "Life is what happens to you while you're busy making other plans.",
		Author: "John Lennon",
	},
	{
		Text:   "To know oneself is to study oneself in action with another person.",
		Author: "Bruce Lee",
	},
	{
		Text:   "Many people have gone further than they thought they could because someone else thought they could.",
		Author: "Someone",
	},
	{
		Text:   "We must become the change we want to see.",
		Author: "Mahatma Gandhi",
	},
	{
		Text:   "Happiness comes when your work and words are of benefit to yourself and others.",
		Author: "Buddha",
	},
	{
		Text:   "If the shoe doesn't fit, must we change the foot?",
		Author: "Gloria Steinem",
	},
	{
		Text:   "Real magic in relationships means an absence of judgement of others.",
		Author: "Wayne Dyer",
	},
	{
		Text:   "Arriving at one point is the starting point to another.",
		Author: "John Dewey",
	},
	{
		Text:   "I begin with an idea and then it becomes something else.",
		Author: "Pablo Picasso",
	},
	{
		Text:   "The weak can never forgive. Forgiveness is the attribute of the strong.",
		Author: "Mohandas Gandhi",
	},
	{
		Text:   "You do not become good by trying to be good, but by finding the goodness that is already within you.",
		Author: "Eckhart Tolle",
	},
	{
		Text:   "Try not to become a man of success but rather try to become a man of value.",
		Author: "Albert Einstein",
	},
	{
		Text:   "We are all something, but none of us are everything.",
		Author: "Blaise Pascal",
	},
	{Text: "What we think, we become.", Author: "Buddha"},
	{
		Text:   "Every person, all the events of your life are there because you have drawn them there. What you choose to do with them is up to you.",
		Author: "Richard Bach",
	},
	{
		Text:   "If there is no struggle, there is no progress.",
		Author: "Frederick Douglass",
	},
	{
		Text:   "It is not so important to know everything as to appreciate what we learn.",
		Author: "Hannah More",
	},
	{
		Text:   "Joy is what happens to us when we allow ourselves to recognize how good things really are.",
		Author: "Marianne Williamson",
	},
	{
		Text:   "The highest stage in moral ure at which we can arrive is when we recognize that we ought to control our thoughts.",
		Author: "Charles Darwin",
	},
	{
		Text:   "When you dance, your purpose is not to get to a certain place on the floor. It's to enjoy each step along the way.",
		Author: "Wayne Dyer",
	},
	{
		Text:   "Genuine love should first be directed at oneself � if we do not love ourselves, how can we love others?",
		Author: "Dalai Lama",
	},
	{
		Text:   "He that never changes his opinions, never corrects his mistakes, and will never be wiser on the morrow than he is today.",
		Author: "Tryon Edwards",
	},
	{
		Text:   "Promises are the uniquely human way of ordering the future, making it predictable and reliable to the extent that this is humanly possible.",
		Author: "Hannah Arendt",
	},
	{
		Text:   "You can complain because roses have thorns, or you can rejoice because thorns have roses.",
		Author: "Ziggy",
	},
	{
		Text:   "You can't trust without risk but neither can you live in a cocoon.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "It is not so important to know everything as to appreciate what we learn.",
		Author: "Hannah More",
	},
	{
		Text:   "While we stop to think, we often miss our opportunity.",
		Author: "Publilius Syrus",
	},
	{
		Text:   "The most complicated achievements of thought are possible without the assistance of consciousness.",
		Author: "Sigmund Freud",
	},
	{
		Text:   "To ensure good health: eat lightly, breathe deeply, live moderately, cultivate cheerfulness, and maintain an interest in life.",
		Author: "William Londen",
	},
	{
		Text:   "The universe is transformation; our life is what our thoughts make it.",
		Author: "Marcus Aurelius",
	},
	{
		Text:   "The first step to getting the things you want out of life is this: decide what you want.",
		Author: "Ben Stein",
	},
	{
		Text:   "To get something you never had, you have to do something you never did.",
		Author: "Someone",
	},
	{Text: "A good rest is half the work.", Author: "Someone"},
	{
		Text:   "Belief consists in accepting the affirmations of the soul; Unbelief, in denying them.",
		Author: "Ralph Emerson",
	},
	{
		Text:   "There are things so deep and complex that only intuition can reach it in our stage of development as human beings.",
		Author: "John Astin",
	},
	{
		Text:   "Real success is finding your lifework in the work that you love.",
		Author: "David McCullough",
	},
	{
		Text:   "The doors we open and close each day decide the lives we live.",
		Author: "Flora Whittemore",
	},
	{
		Text:   "Though no one can go back and make a brand new start, anyone can start from now and make a brand new ending.",
		Author: "Someone",
	},
	{
		Text:   "Goals are the fuel in the furnace of achievement.",
		Author: "Brian Tracy",
	},
	{
		Text:   "The art of progress is to preserve order amid change, and to preserve change amid order.",
		Author: "Alfred Whitehead",
	},
	{
		Text:   "Tension is who you think you should be. Relaxation is who you are.",
		Author: "Chinese proverb",
	},
	{
		Text:   "Take no thought of who is right or wrong or who is better than. Be not for or against.",
		Author: "Bruce Lee",
	},
	{
		Text:   "To be beautiful means to be yourself. You don�t need to be accepted by others. You need to accept yourself.",
		Author: "Thich Nhat Hanh",
	},
	{
		Text:   "You will never be happy if you continue to search for what happiness consists of. You will never live if you are looking for the meaning of life.",
		Author: "Albert Camus",
	},
	{
		Text:   "All seasons are beautiful for the person who carries happiness within.",
		Author: "Horace Friess",
	},
	{
		Text:   "When you meet someone better than yourself, turn your thoughts to becoming his equal. When you meet someone not as good as you are, look within and examine your own self.",
		Author: "Confucius",
	},
	{
		Text:   "Real magic in relationships means an absence of judgement of others.",
		Author: "Wayne Dyer",
	},
	{
		Text:   "I have often regretted my speech, never my silence.",
		Author: "Publilius Syrus",
	},
	{
		Text:   "I believe that every person is born with talent.",
		Author: "Maya Angelou",
	},
	{
		Text:   "Happiness is not something ready made. It comes from your own actions.",
		Author: "Dalai Lama",
	},
	{
		Text:   "From little acorns mighty oaks do grow.",
		Author: "American proverb",
	},
	{
		Text:   "We are shaped by our thoughts; we become what we think. When the mind is pure, joy follows like a shadow that never leaves.",
		Author: "Buddha",
	},
	{
		Text:   "Inspiration exists, but it has to find us working.",
		Author: "Pablo Picasso",
	},
	{Text: "You only lose what you cling to.", Author: "Buddha"},
	{Text: "A jug fills drop by drop.", Author: "Buddha"},
	{
		Text:   "A thing long expected takes the form of the unexpected when at last it comes.",
		Author: "Mark Twain",
	},
	{
		Text:   "It is better to understand a little than to misunderstand a lot.",
		Author: "Anatole France",
	},
	{
		Text:   "The cause is hidden. The effect is visible to all.",
		Author: "Ovid",
	},
	{
		Text:   "What is new in the world? Nothing. What is old in the world? Nothing. Everything has always been and will always be.",
		Author: "Sai Baba",
	},
	{
		Text:   "Almost everything comes from nothing.",
		Author: "Henri Amiel",
	},
	{
		Text:   "I'm not afraid of storms, for Im learning how to sail my ship.",
		Author: "Louisa Alcott",
	},
	{
		Text:   "If you correct your mind, the rest of your life will fall into place.",
		Author: "Lao Tzu",
	},
	{
		Text:   "There is only one success � to be able to spend your life in your own way.",
		Author: "Christopher Morley",
	},
	{
		Text:   "Spirituality can be severed from both vicious sectarianism and thoughtless banalities. Spirituality, I have come to see, is nothing less than the thoughtful love of life.",
		Author: "Robert C. Solomon",
	},
	{Text: "Wisdom begins in wonder.", Author: "Socrates"},
	{
		Text:   "We must not allow ourselves to become like the system we oppose.",
		Author: "Bishop Desmond Tutu",
	},
	{
		Text:   "Most powerful is he who has himself in his own power.",
		Author: "Seneca",
	},
	{
		Text:   "I don't believe in failure. It is not failure if you enjoyed the process.",
		Author: "Oprah Winfrey",
	},
	{
		Text:   "It takes both sunshine and rain to make a rainbow.",
		Author: "Someone",
	},
	{
		Text:   "Sometimes our fate resembles a fruit tree in winter. Who would think that those branches would turn green again and blossom, but we hope it, we know it.",
		Author: "Johann Wolfgang von Goethe",
	},
	{
		Text:   "I am not bothered by the fact that I am unknown. I am bothered when I do not know others.",
		Author: "Confucius",
	},
	{
		Text:   "Logic will get you from A to B. Imagination will take you everywhere.",
		Author: "Albert Einstein",
	},
	{Text: "A stumble may prevent a fall.", Author: "Someone"},
	{
		Text:   "Try not to become a man of success but rather try to become a man of value.",
		Author: "Albert Einstein",
	},
	{
		Text:   "A rolling stone gathers no moss.",
		Author: "Publilius Syrus",
	},
	{
		Text:   "The greatest mistake you can make in life is to be continually fearing you will make one.",
		Author: "Elbert Hubbard",
	},
	{
		Text:   "If you want a thing done well, do it yourself.",
		Author: "Napoleon Bonaparte",
	},
	{
		Text:   "Quality is never an accident; it is always the result of intelligent effort.",
		Author: "John Ruskin",
	},
	{
		Text:   "To study and not think is a waste. To think and not study is dangerous.",
		Author: "Confucius",
	},
	{
		Text:   "Do not expect the world to look bright, if you habitually wear gray-brown glasses.",
		Author: "Tomas Eliot",
	},
	{
		Text:   "The years teach much which the days never know.",
		Author: "Ralph Emerson",
	},
	{
		Text:   "You're not obligated to win. You're obligated to keep trying to do the best you can every day.",
		Author: "Marian Edelman",
	},
	{
		Text:   "Man cannot discover new oceans unless he has the courage to lose sight of the shore.",
		Author: "Andr� Gide",
	},
	{
		Text:   "Trust your hunches. They're usually based on facts filed away just below the conscious level.",
		Author: "Joyce Brothers",
	},
	{
		Text:   "We must become the change we want to see.",
		Author: "Mahatma Gandhi",
	},
	{
		Text:   "Many of life's failures are people who did not realize how close they were to success when they gave up.",
		Author: "Thomas Edison",
	},
	{
		Text:   "Character cannot be developed in ease and quiet. Only through experience of trial and suffering can the soul be strengthened, vision cleared, ambition inspired, and success achieved.",
		Author: "Helen Keller",
	},
	{
		Text:   "Setting an example is not the main means of influencing another, it is the only means.",
		Author: "Albert Einstein",
	},
	{
		Text:   "Some people think it's holding that makes one strong � sometimes it's letting go.",
		Author: "Someone",
	},
	{
		Text:   "Action will remove the doubts that theory cannot solve.",
		Author: "Tehyi Hsieh",
	},
	{
		Text:   "An optimist is a person who sees a green light everywhere, while the pessimist sees only the red spotlight... The truly wise person is colour-blind.",
		Author: "Albert Schweitzer",
	},
	{
		Text:   "Success means having the courage, the determination, and the will to become the person you believe you were meant to be.",
		Author: "George Sheehan",
	},
	{
		Text:   "And as we let our own light shine, we unconsciously give other people permission to do the same.",
		Author: "Nelson Mandela",
	},
	{
		Text:   "Before you put on a frown, make absolutely sure there are no smiles available.",
		Author: "Jim Beggs",
	},
	{
		Text:   "Success means having the courage, the determination, and the will to become the person you believe you were meant to be.",
		Author: "George Sheehan",
	},
	{
		Text:   "Nothing is predestined: The obstacles of your past can become the gateways that lead to new beginnings.",
		Author: "Ralph Blum",
	},
	{
		Text:   "For every failure, there's an alternative course of action. You just have to find it. When you come to a roadblock, take a detour.",
		Author: "Mary Kay Ash",
	},
	{
		Text:   "Change your thoughts, change your life!",
		Author: "Someone",
	},
	{
		Text:   "You don't drown by falling in water. You drown by staying there.",
		Author: "Someone",
	},
	{
		Text:   "You, yourself, as much as anybody in the entire universe, deserve your love and affection.",
		Author: "Buddha",
	},
	{
		Text:   "What is necessary to change a person is to change his awareness of himself.",
		Author: "Abraham Maslow",
	},
	{
		Text:   "Adversity causes some men to break, others to break records.",
		Author: "William Ward",
	},
	{
		Text:   "Sometimes it is better to lose and do the right thing than to win and do the wrong thing.",
		Author: "Tony Blair",
	},
	{
		Text:   "We have two ears and one mouth so that we can listen twice as much as we speak.",
		Author: "Epictetus",
	},
	{
		Text:   "Good people are good because they've come to wisdom through failure. We get very little wisdom from success, you know.",
		Author: "William Saroyan",
	},
	{Text: "Fortune befriends the bold.", Author: "John Dryden"},
	{
		Text:   "Everyone has been made for some particular work, and the desire for that work has been put in every heart.",
		Author: "Rumi",
	},
	{
		Text:   "Meaning is not what you start with but what you end up with.",
		Author: "Peter Elbow",
	},
	{
		Text:   "Try not to become a man of success, but rather try to become a man of value.",
		Author: "Albert Einstein",
	},
	{
		Text:   "Always be smarter than the people who hire you.",
		Author: "Lena Horne",
	},
	{
		Text:   "Holding on to anger is like grasping a hot coal with the intent of throwing it at someone else; you are the one who gets burned.",
		Author: "Buddha",
	},
	{
		Text:   "In the end we retain from our studies only that which we practically apply.",
		Author: "Johann Wolfgang von Goethe",
	},
	{
		Text:   "Holding on to anger is like grasping a hot coal with the intent of throwing it at someone else; you are the one who gets burned.",
		Author: "Buddha",
	},
	{
		Text:   "We cannot solve our problems with the same thinking we used when we created them.",
		Author: "Albert Einstein",
	},
	{
		Text:   "Never ignore a gut feeling, but never believe that it's enough.",
		Author: "Robert Heller",
	},
	{
		Text:   "I cannot make my days longer so I strive to make them better.",
		Author: "Henry David Thoreau",
	},
	{
		Text:   "The energy of the mind is the essence of life.",
		Author: "Aristotle",
	},
	{
		Text:   "A person who never made a mistake never tried anything new.",
		Author: "Albert Einstein",
	},
	{
		Text:   "If you must tell me your opinions, tell me what you believe in. I have plenty of douts of my own.",
		Author: "Johann Wolfgang von Goethe",
	},
	{
		Text:   "Your outlook on life is a direct reflection on how much you like yourself.",
		Author: "Lululemon",
	},
	{
		Text:   "Look forward to spring as a time when you can start to see what nature has to offer once again.",
		Author: "Byron Pulsifer",
	},
	{
		Text:   "When performance exceeds ambition, the overlap is called success.",
		Author: "Cullen Hightower",
	},
	{
		Text:   "Never mistake activity for achievement.",
		Author: "John Wooden",
	},
	{
		Text:   "As the rest of the world is walking out the door, your best friends are the ones walking in.",
		Author: "Someone",
	},
	{
		Text:   "Friends are those rare people who ask how we are and then wait to hear the answer.",
		Author: "Ed Cunningham",
	},
	{
		Text:   "He who conquers others is strong; He who conquers himself is mighty.",
		Author: "Lao Tzu",
	},
	{
		Text:   "There are two kinds of failures: those who thought and never did, and those who did and never thought.",
		Author: "Laurence J. Peter",
	},
	{
		Text:   "Adversity has the effect of eliciting talents, which in prosperous circumstances would have lain dormant.",
		Author: "Horace",
	},
	{
		Text:   "Each time we face a fear, we gain strength, courage, and confidence in the doing.",
		Author: "Someone",
	},
	{
		Text:   "Be a good listener. Your ears will never get you in trouble.",
		Author: "Frank Tyger",
	},
	{
		Text:   "I have been impressed with the urgency of doing. Knowing is not enough; we must apply. Being willing is not enough; we must do.",
		Author: "Leonardo da Vinci",
	},
	{
		Text:   "Everyone smiles in the same language.",
		Author: "Someone",
	},
	{
		Text:   "Risk more than others think is safe. Care more than others think is wise. Dream more than others think is practical.Expect more than others think is possible.",
		Author: "Cadet Maxim",
	},
	{
		Text:   "Chaos and Order are not enemies, only opposites.",
		Author: "Richard Garriott",
	},
	{
		Text:   "Mistakes are always forgivable, if one has the courage to admit them.",
		Author: "Bruce Lee",
	},
	{
		Text:   "Friendship isn't a big thing. It's a million little things.",
		Author: "Someone",
	},
	{
		Text:   "Everything is perfect in the universe � even your desire to improve it.",
		Author: "Wayne Dyer",
	},
	{
		Text:   "The road leading to a goal does not separate you from the destination; it is essentially a part of it.",
		Author: "Charles DeLint",
	},
}
