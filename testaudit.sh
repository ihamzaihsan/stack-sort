#!/bin/bash
# Manual audit examples; inspect the output of each command.
./push-swap
./push-swap "2 1 3 6 5 8"
./push-swap "0 1 2 3 4 5"
./push-swap "0 one 2 3"
./push-swap "1 2 2 3"
./push-swap "427 182 941 308 729"
./push-swap "136 753 596 250 131"
./checker
./checker "0 one 2 3"
echo -e "sa\npb\nrrr" | ./checker "0 9 1 8 2 7 3 6 4 5"
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa" | ./checker "0 9 1 8 2"
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
ARG="757 303 139 343 675 919 793 119 916 513 706 312 712 656 506 173 491 791 852 436 836 871 394 884 351 180 693 406 385 67 553 945 877 854 775 982 182 439 306 531 927 648 10 481 148 165 702 352 942 94 777 890 230 391 810 536 714 655 7 390 611 816 443 613 328 857 336 309 225 28 318 593 311 564 296 359 760 938 662 872 603 473 437 883 740 12 44 541 289 749 873 342 460 413 501 654 243 112 71 49"; ./push-swap "$ARG" | ./checker "$ARG"
