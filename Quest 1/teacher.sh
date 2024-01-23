export INTERVIEW_NO=$(head -n 179 */Buckingham_Place | tail -n1 | sed 's/[^0-9]//g')

echo $INTERVIEW_NO

cat */*-$INTERVIEW_NO

echo $MAIN_SUSPECT