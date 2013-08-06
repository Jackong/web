/**
 * User: Jackong
 * Date: 13-8-6
 * Time: 下午9:20
 */
package i

type InputError string

func (this InputError) Error() string {
	return string(this)
}
