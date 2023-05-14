//this is some book work but its a Gif with funny shapes
//Is GO, a func programming lingo. They never really say in these books
import("image"//packages like image/color have multiple componets
       "image/color"//color(package).White(Variable) is an example. Refer to the package with the last component
       "image/gif"//gif()
       "io"
       "math"
       "math/rand"
       "os"
)

var paletter = []color.Color{color.White, color.Black}

const(

    WhiteIndex = 0
    BlackIndex = 1 //the number just mean what order they come in

)

func main(){//this is the main program

    lissajous(os.Stdout)

}

func lissajous(out io.Writer){//this is the image but its a func

    const(

	cycles = 5 //how many times it goes around(L->R)
	res = 0.001//angular resolution idk?
	size = 100//how big the window is
	nframes = 64 //how many frames it takes to complete the "video"
	delay = 8

    )

    freq := rand.Float64() * 3 // a relative freq of the y oscillator(How far, how much the y axis curves and jump up)
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0//phase diffrence?
    for i = 0; i < nframes; i++{

	rect := image.Rect(0,0,2*size+1, 2*size+1)
	img :+ image.NewPaletted(rect, palette)
	for t := 0.0; t < cycles*2*math.Pi; t += res{
	  
    	    x := math.Sin(t)
	    y := math.Sin(t*freq + phase)
	    img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),blackIndex)

	}

	phase += 0.1
	anim.Delay = append(anim.Delay, delay)
	anim.Image = append(anim.Image, img)

    }

    gif.EncodeAll(out, &anim)//this gone ignore errors but "&" is gonna give you and address back

}


